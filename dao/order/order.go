package order

import (
	"database/sql"
	"miaosha/internal/code"
	"miaosha/internal/db"
	"miaosha/model"
)

type Dao struct {
	db *sql.DB
}

func New() *Dao {
	return &Dao{}
}

var (
	_countSql     = `select count(1) from miaosha_order where user_id = ? and goods_id = ?`
	_decrStockSql = `update miaosha_goods set stock = stock - 1 where id = ? and stock > 0`
	_insertSql    = `insert into miaosha_order(id, user_id, goods_id, create_time, status) values(?, ?, ?, ?, ?)`
)

func (d *Dao) Count(userId, goodsId int64) (count int64, err error) {
	var (
		stmt *sql.Stmt
	)
	if stmt, err = db.Conn().Prepare(_countSql); err != nil {
		return
	}
	if err = stmt.QueryRow(userId, goodsId).Scan(&count); err != nil {
		return
	}
	return
}

func (d *Dao) Miaosha(order *model.Order) (err error) {
	var (
		tx        *sql.Tx
		rs, rs1   sql.Result
		aff, aff1 int64
	)
	if tx, err = db.Conn().Begin(); err != nil {
		return
	}
	defer tx.Rollback()
	// 减库存
	if rs, err = tx.Exec(_decrStockSql, order.GoodsId); err != nil {
		return
	}
	if aff, err = rs.RowsAffected(); err != nil {
		return
	}
	// 生成订单
	if rs1, err = tx.Exec(_insertSql, order.Id, order.UserId, order.GoodsId, order.CreateTime, order.Status); err != nil {
		return
	}
	if aff1, err = rs1.RowsAffected(); err != nil {
		return
	}
	if aff == 0 || aff1 == 0 {
		err = code.MiaoshaFailed
		return
	}
	err = tx.Commit()
	return
}
