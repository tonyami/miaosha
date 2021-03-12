package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"miaosha/infra/db"
	"miaosha/model"
)

func GetOrderList(userId int64, status string, page int) (list []model.Order, err error) {
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
	sqlStr := "select * from miaosha_order where user_id = ?"
	if err = db.Conn().Select(&list, sqlStr+s, args...); err != nil {
		log.Printf("GetOrderList() failed, err: %v", err)
	}
	return
}

func GetOrderById(id int64) (order model.Order, err error) {
	sqlStr := "select * from miaosha_order where id = ?"
	if err = db.Conn().Get(&order, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Printf("GetOrderById() failed, err: %v", err)
		}
	}
	return
}

func GetOrderByOrderId(orderId string) (order model.Order, err error) {
	sqlStr := "select * from miaosha_order where order_id = ?"
	if err = db.Conn().Get(&order, sqlStr, orderId); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Printf("GetOrderByOrderId() failed, err: %v", err)
		}
	}
	return
}

func CreateOrder(order model.Order) (err error) {
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
	sqlStr := "insert into miaosha_order(order_id, user_id, goods_id, goods_name, goods_img, goods_price, status) values(?, ?, ?, ?, ?, ?, ?)"
	if rs, err = tx.Exec(sqlStr, order.OrderId, order.UserId, order.GoodsId, order.GoodsName, order.GoodsImg, order.GoodsPrice, order.Status); err != nil {
		log.Printf("CreateOrder() failed, err: %v, order: %v", err, order)
		return
	}
	if rows, err = rs.RowsAffected(); err != nil || rows == 0 {
		err = errors.New("创建订单失败")
		return
	}
	err = tx.Commit()
	return
}

func CountRepeatableOrder(userId, goodsId int64) (count int64, err error) {
	SqlStr := "select count(*) from miaosha_order where user_id = ? and goods_id = ? and status != ?"
	if err = db.Conn().Get(&count, SqlStr, userId, goodsId, model.Closed); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Printf("GetOrderById() failed, err: %v", err)
		}
	}
	return
}

func CountOrder(userId int64) (count model.OrderCount, err error) {
	sqlStr := "select ifnull(sum(case when status = 0 or status = 1 then 1 else 0 end), 0) 'unfinished', ifnull(sum(case when status = 2 then 1 else 0 end), 0) 'finished', ifnull(sum(case when status = -1 then 1 else 0 end), 0) 'closed' from `miaosha_order` where `user_id` = ?"
	if err = db.Conn().Get(&count, sqlStr, userId); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Printf("stmt.QueryRow(_countByStatusSql, %d).Scan() failed, err: %v", userId, err)
		}
	}
	return
}

func CloseOrder(orderId string, goodsId int64) (err error) {
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
	if rs, err = tx.Exec(incrStockSql, goodsId); err != nil {
		log.Printf("incrStock failed, err: %v, goodsId: %v", err, goodsId)
		return
	}
	if rows, err = rs.RowsAffected(); err != nil || rows == 0 {
		err = errors.New("加库存失败")
		return
	}
	// 修改订单状态
	sqlStr := "update miaosha_order set status = ? where order_id = ? and status = ?"
	if rs, err = tx.Exec(sqlStr, model.Closed, orderId, model.Unpaid); err != nil {
		log.Printf("UpdateOrderStatus() failed, err: %v", err)
		return
	}
	if rows, err = rs.RowsAffected(); err != nil || rows == 0 {
		err = errors.New("修改订单状态失败")
		return
	}
	err = tx.Commit()
	return
}
