package goods

import (
	"database/sql"
	"miaosha/conf"
	"miaosha/internal/db"
	"miaosha/model"
)

type Dao struct {
	db *sql.DB
}

func New() *Dao {
	return &Dao{
		db: db.New(conf.Conf.DB),
	}
}

var (
	_getListSql = `select id, name, img, origin_price, price, stock, start_time, end_time
					from miaosha_goods order by id desc limit ?, ?`
	_getSql = `select id, name, img, origin_price, price, stock, start_time, end_time
				from miaosha_goods where id = ? limit 1`
)

func (d *Dao) GetList(page, size int) (list []*model.Goods, err error) {
	stmt, err := d.db.Prepare(_getListSql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query((page-1)*size, size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list = []*model.Goods{}
	for rows.Next() {
		goods := &model.Goods{}
		if err = rows.Scan(&goods.Id, &goods.Name, &goods.Img, &goods.OriginPrice, &goods.Price, &goods.Stock, &goods.StartTime, &goods.EndTime); err != nil {
			return
		}
		list = append(list, goods)
	}
	return
}

func (d *Dao) Get(goodsId int64) (goods *model.Goods, err error) {
	stmt, err := d.db.Prepare(_getSql)
	if err != nil {
		return
	}
	defer stmt.Close()
	goods = &model.Goods{}
	if err = stmt.QueryRow(goodsId).Scan(&goods.Id, &goods.Name, &goods.Img, &goods.OriginPrice, &goods.Price, &goods.Stock, &goods.StartTime, &goods.EndTime); err != nil {
		goods = nil
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	return
}
