package order

import (
	"database/sql"
	"miaosha/internal/code"
	"miaosha/internal/db"
	"miaosha/model"
)

type Dao struct{}

func New() *Dao {
	return &Dao{}
}

var (
	_getSql       = `select o.id, o.user_id, o.goods_id, g.name, g.img, g.price, o.create_time, o.status from miaosha_order o left join miaosha_goods g on o.goods_id = g.id where o.id = ? limit 1`
	_getListSql   = `select o.id, o.goods_id, g.name, g.img, g.price, o.create_time, o.status from miaosha_order o left join miaosha_goods g on o.goods_id = g.id where o.user_id = ? order by o.create_time desc limit ?, ?`
	_getList2Sql  = `select o.id, o.goods_id, g.name, g.img, g.price, o.create_time, o.status from miaosha_order o left join miaosha_goods g on o.goods_id = g.id where o.user_id = ? and o.status = ? order by o.create_time desc limit ?, ?`
	_countSql     = `select count(1) from miaosha_order where user_id = ? and goods_id = ?`
	_decrStockSql = `update miaosha_goods set stock = stock - 1 where id = ? and stock > 0`
	_insertSql    = `insert into miaosha_order(id, user_id, goods_id, create_time, status) values(?, ?, ?, ?, ?)`
)

func (*Dao) Get(id string) (order *model.OrderDTO, err error) {
	var stmt *sql.Stmt
	if stmt, err = db.Conn().Prepare(_getSql); err != nil {
		return
	}
	order = &model.OrderDTO{}
	if err = stmt.QueryRow(id).Scan(&order.Id, &order.UserId, &order.GoodsId, &order.Name, &order.Img, &order.Price, &order.CreateTime, &order.Status); err != nil {
		order = nil
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	return
}

func (*Dao) GetList(userId int64, page, size int, status string) (orders []*model.OrderDTO, err error) {
	var (
		stmt *sql.Stmt
		rows *sql.Rows
	)
	if len(status) == 0 {
		if stmt, err = db.Conn().Prepare(_getListSql); err != nil {
			return
		}
		defer stmt.Close()
		if rows, err = stmt.Query(userId, (page-1)*size, size); err != nil {
			return
		}
		defer rows.Close()
	} else {
		if stmt, err = db.Conn().Prepare(_getList2Sql); err != nil {
			return
		}
		defer stmt.Close()
		if rows, err = stmt.Query(userId, status, (page-1)*size, size); err != nil {
			return
		}
		defer rows.Close()
	}
	orders = []*model.OrderDTO{}
	for rows.Next() {
		order := model.OrderDTO{}
		if err = rows.Scan(&order.Id, &order.GoodsId, &order.Name, &order.Img, &order.Price, &order.CreateTime, &order.Status); err != nil {
			return
		}
		orders = append(orders, &order)
	}
	return
}

func (*Dao) Count(userId, goodsId int64) (count int64, err error) {
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

func (*Dao) Miaosha(order *model.Order) (err error) {
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
