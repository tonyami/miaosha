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
	_getListSql = `select m.id miaosha_id, m.goods_id, m.miaosha_price, m.miaosha_stock, m.start_time, m.end_time,
						g.id, g.goods_name, g.goods_img, g.goods_price
					from miaosha_goods m left join goods g on m.goods_id = g.id
					order by m.id desc limit ?, ?`
	_getSql = `select m.id miaosha_id, m.goods_id, m.miaosha_price, m.miaosha_stock, m.start_time, m.end_time,
					g.id, g.goods_name, g.goods_img, g.goods_price
				from miaosha_goods m left join goods g on m.goods_id = g.id
				where m.id = ? limit 1`
)

func (d *Dao) GetList(page, size int) (list []*model.GoodsDTO, err error) {
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
	list = []*model.GoodsDTO{}
	for rows.Next() {
		miaoshaGoods := &model.MiaoshaGoods{}
		goods := &model.Goods{}
		if err = rows.Scan(&miaoshaGoods.Id, &miaoshaGoods.GoodsId, &miaoshaGoods.MiaoshaPrice, &miaoshaGoods.MiaoshaStock, &miaoshaGoods.StartTime, &miaoshaGoods.EndTime,
			&goods.Id, &goods.GoodsName, &goods.GoodsImg, &goods.GoodsPrice); err != nil {
			return
		}
		dto := &model.GoodsDTO{
			MiaoshaGoods: miaoshaGoods,
			Goods:        goods,
		}
		list = append(list, dto)
	}
	return
}

func (d *Dao) Get(goodsId int64) (goodsDTO *model.GoodsDTO, err error) {
	stmt, err := d.db.Prepare(_getSql)
	if err != nil {
		return
	}
	defer stmt.Close()
	miaoshaGoods := &model.MiaoshaGoods{}
	goods := &model.Goods{}
	if err = stmt.QueryRow(goodsId).Scan(&miaoshaGoods.Id, &miaoshaGoods.GoodsId, &miaoshaGoods.MiaoshaPrice, &miaoshaGoods.MiaoshaStock, &miaoshaGoods.StartTime, &miaoshaGoods.EndTime,
		&goods.Id, &goods.GoodsName, &goods.GoodsImg, &goods.GoodsPrice); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	goodsDTO = &model.GoodsDTO{
		MiaoshaGoods: miaoshaGoods,
		Goods:        goods,
	}
	return
}
