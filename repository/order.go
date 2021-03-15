package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"miaosha/infra/db"
	"miaosha/model"
)

func GetOrderList(userId int64, status string, page int) (list []model.OrderInfo, err error) {
	args := []interface{}{userId}
	s := " "
	w := "and"
	if status == "unfinished" {
		s += fmt.Sprintf("%s (status = ? or status = ?)", w)
		args = append(args, model.Unpaid, model.Paying)
		w = "and"
	} else if status == "finished" {
		s += fmt.Sprintf("%s status = ?", w)
		args = append(args, model.Paid)
		w = "and"
	} else if status == "closed" {
		s += fmt.Sprintf("%s status = ?", w)
		args = append(args, model.Closed)
		w = "and"
	}
	s += " order by id desc limit ?, ?"
	args = append(args, (page-1)*PageSize, PageSize)
	sqlStr := "select * from miaosha_order_info where user_id = ?"
	if err = db.Conn().Select(&list, sqlStr+s, args...); err != nil {
		log.Printf("GetOrderList() failed, err: %v", err)
	}
	return
}

func GetOrderByOrderId(orderId string) (order model.OrderInfo, err error) {
	sqlStr := "select * from miaosha_order_info where order_id = ?"
	if err = db.Conn().Get(&order, sqlStr, orderId); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Printf("GetOrderByOrderId() failed, err: %v", err)
		}
	}
	return
}

func CreateOrder(order model.OrderInfo) (err error) {
	tx, err := db.Conn().Begin()
	if err != nil {
		log.Printf("db.Conn().Begin() failed, err: %v", err)
		return
	}
	defer tx.Rollback()
	var (
		rs   sql.Result
		rows int64
	)
	// 减库存
	decrStockSql := "update miaosha_goods set stock = stock - 1 where id = ? and stock > 0"
	if rs, err = tx.Exec(decrStockSql, order.GoodsId); err != nil {
		log.Printf("decrStock failed, err: %v, goodsId: %v", err, order.GoodsId)
		return
	}
	if rows, err = rs.RowsAffected(); err != nil || rows == 0 {
		err = errors.New("减库存失败")
		return
	}
	// 创建订单
	orderStr := "insert into miaosha_order(order_id, user_id, goods_id) values(?, ?, ?)"
	if rs, err = tx.Exec(orderStr, order.OrderId, order.UserId, order.GoodsId); err != nil {
		log.Printf("CreateOrder() failed, err: %v, order: %v", err, order)
		return
	}
	if rows, err = rs.RowsAffected(); err != nil || rows == 0 {
		err = errors.New("创建订单失败")
		return
	}
	// 创建订单信息
	orderInfoStr := "insert into miaosha_order_info(order_id, user_id, goods_id, goods_name, goods_img, goods_price, status) values(?, ?, ?, ?, ?, ?, ?)"
	if rs, err = tx.Exec(orderInfoStr, order.OrderId, order.UserId, order.GoodsId, order.GoodsName, order.GoodsImg, order.GoodsPrice, order.Status); err != nil {
		log.Printf("CreateOrderInfo() failed, err: %v, order: %v", err, order)
		return
	}
	if rows, err = rs.RowsAffected(); err != nil || rows == 0 {
		err = errors.New("创建订单信息失败")
		return
	}
	err = tx.Commit()
	return
}

func CountOrderByUidAndGid(userId, goodsId int64) (count int64, err error) {
	sqlStr := "select count(*) from miaosha_order_info where user_id = ? and goods_id = ? and status != ?"
	if err = db.Conn().Get(&count, sqlStr, userId, goodsId, model.Closed); err != nil {
		log.Printf("db.Get() failed, err: %v", err)
	}
	return
}

func CountOrder(userId int64) (count model.OrderCount, err error) {
	sqlStr := "select ifnull(sum(case when status = 0 or status = 1 then 1 else 0 end), 0) 'unfinished', ifnull(sum(case when status = 2 then 1 else 0 end), 0) 'finished', ifnull(sum(case when status = -1 then 1 else 0 end), 0) 'closed' from miaosha_order_info where user_id = ?"
	if err = db.Conn().Get(&count, sqlStr, userId); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Printf("db.Get() failed, err: %v", err)
		}
	}
	return
}

func CloseOrder(order model.OrderInfo) (err error) {
	tx, err := db.Conn().Begin()
	if err != nil {
		log.Printf("db.Conn().Begin() failed, err: %v", err)
		return
	}
	defer tx.Rollback()
	var (
		rs   sql.Result
		rows int64
	)
	// 加库存
	incrStockSql := "update miaosha_goods set stock = stock + 1 where id = ?"
	if rs, err = tx.Exec(incrStockSql, order.GoodsId); err != nil {
		log.Printf("incrStock failed, err: %v, goodsId: %v", err, order.GoodsId)
		return
	}
	if rows, err = rs.RowsAffected(); err != nil || rows == 0 {
		err = errors.New("加库存失败")
		return
	}
	// 删除订单
	deleteStr := "delete from miaosha_order where order_id = ?"
	if rs, err = tx.Exec(deleteStr, order.OrderId); err != nil {
		log.Printf("DeleteOrder() failed, err: %v", err)
		return
	}
	if rows, err = rs.RowsAffected(); err != nil || rows == 0 {
		err = errors.New("删除订单失败")
		return
	}
	// 修改订单信息状态
	updateStr := "update miaosha_order_info set status = ? where order_id = ? and status = ?"
	if rs, err = tx.Exec(updateStr, model.Closed, order.OrderId, model.Unpaid); err != nil {
		log.Printf("UpdateOrderStatus() failed, err: %v", err)
		return
	}
	if rows, err = rs.RowsAffected(); err != nil || rows == 0 {
		err = errors.New("修改订单信息状态失败")
		return
	}
	err = tx.Commit()
	return
}
