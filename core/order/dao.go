package order

import (
	"database/sql"
	"log"
	"miaosha/service"
)

type Dao struct {
	db *sql.DB
}

func NewDao(db *sql.DB) *Dao {
	return &Dao{
		db: db,
	}
}

var (
	_countPurchasedSql = "select count(*) from `miaosha_order` where `user_id` = ? and `goods_id` = ? and `status` != ?"
	_insertSql         = "insert into `miaosha_order`(`user_id`, `goods_id`, `goods_name`, `goods_img`, `goods_price`) values(?, ?, ?, ?, ?)"
	_getSql            = "select `id`, `user_id`, `goods_id`, `goods_name`, `goods_img`, `goods_price`, `status`, `create_time`, `update_time` from `miaosha_order` where `id` = ? and `user_id` = ? limit 1"
	_getListSql        = "select `id`, `user_id`, `goods_id`, `goods_name`, `goods_img`, `goods_price`, `status`, `create_time`, `update_time` from `miaosha_order` where `user_id` = ?"
	_closeSql          = "update `miaosha_order` set `status` = ? where `id` = ? and `status` = ?"
)

func (dao *Dao) CountPurchased(userId, goodsId int64) (count int64, err error) {
	if err = dao.db.QueryRow(_countPurchasedSql, userId, goodsId, service.Closed).Scan(&count); err != nil {
		log.Printf("dao.db.QueryRow(_countPurchasedSql) failed, userId: %d, goodsId: %d, err: %v", userId, goodsId, err)
		return
	}
	return
}

func (dao *Dao) Insert(order *Order) (orderId int64, err error) {
	var rs sql.Result
	if rs, err = dao.db.Exec(_insertSql, order.UserId, order.GoodsId, order.GoodsName, order.GoodsImg, order.GoodsPrice); err != nil {
		log.Printf("dao.db.Exec(_insertSql) failed, order: %#v, err: %v", order, err)
		return
	}
	if orderId, err = rs.LastInsertId(); err != nil {
		return
	}
	return
}

func (dao *Dao) Get(id, userId int64) (order *Order, err error) {
	order = &Order{}
	if err = dao.db.QueryRow(_getSql, id, userId).Scan(&order.Id, &order.UserId, &order.GoodsId, &order.GoodsName, &order.GoodsImg, &order.GoodsPrice, &order.Status, &order.CreateTime, &order.UpdateTime); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Printf("stmt.QueryRow(%d).Scan() failed, err: %v", id, err)
		}
	}
	return
}

func (dao *Dao) GetList(userId int64, page, size int, status string) (list []*Order, err error) {
	var rows *sql.Rows
	_sql := _getListSql
	if len(status) > 0 {
		_sql += " and `status` = ?"
		_sql += " order by id desc limit ?, ?"
		if rows, err = dao.db.Query(_sql, userId, status, (page-1)*size, size); err != nil {
			log.Printf("dao.db.Query(_getListSql, %d, %s, %d, %d) failed, err: %v", userId, status, (page-1)*size, size, err)
			return
		}
	} else {
		_sql += " order by id desc limit ?, ?"
		if rows, err = dao.db.Query(_sql, userId, (page-1)*size, size); err != nil {
			log.Printf("dao.db.Query(_getListSql, %d, %d, %d) failed, err: %v", userId, (page-1)*size, size, err)
			return
		}
	}
	defer rows.Close()
	list = []*Order{}
	for rows.Next() {
		order := Order{}
		if err = rows.Scan(&order.Id, &order.UserId, &order.GoodsId, &order.GoodsName, &order.GoodsImg, &order.GoodsPrice, &order.Status, &order.CreateTime, &order.UpdateTime); err != nil {
			log.Printf("rows.Scan() failed, err: %v", err)
			return
		}
		list = append(list, &order)
	}
	return
}

func (dao *Dao) Close(id int64) (err error) {
	var rs sql.Result
	if rs, err = dao.db.Exec(_closeSql, service.Closed, id, service.Unpaid); err != nil {
		log.Printf("dao.db.Exec(_closeSql, %d) failed, err: %v", id, err)
		return
	}
	if _, err = rs.RowsAffected(); err != nil {
		log.Printf("rs.RowsAffected() failed, err: %v", err)
	}
	return
}
