package goods

import (
	"database/sql"
	"miaosha/conf/mysql"
	"miaosha/model"
)

type Dao struct {
	db *sql.DB
}

func New() *Dao {
	return &Dao{
		db: mysql.New(),
	}
}

var (
	_queryAllSql = `SELECT
						m.id miaosha_id,
						m.goods_id,
						m.miaosha_price,
						m.miaosha_stock,
						m.start_time,
						m.end_time,
						g.id,
						g.goods_name,
						g.goods_img,
						g.goods_price
					FROM
						miaosha_goods m
					LEFT JOIN goods g ON m.goods_id = g.id
					ORDER BY m.id DESC 
					LIMIT ?, ?`
	_queryByIdSql = `SELECT
						m.id miaosha_id,
						m.goods_id,
						m.miaosha_price,
						m.miaosha_stock,
						m.start_time,
						m.end_time,
						g.id,
						g.goods_name,
						g.goods_img,
						g.goods_price
					FROM
						miaosha_goods m
					LEFT JOIN goods g ON m.goods_id = g.id
					WHERE m.id = ?
					LIMIT 1`
)

func (d *Dao) QueryAll(page, size int) (goodsList []*model.MiaoshaGoodsDTO, err error) {
	stmt, err := d.db.Prepare(_queryAllSql)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query((page-1)*size, size)
	if err != nil {
		return
	}
	defer rows.Close()
	goodsList = []*model.MiaoshaGoodsDTO{}
	for rows.Next() {
		miaoshaGoods := model.MiaoshaGoods{}
		goods := model.Goods{}
		if err = rows.Scan(&miaoshaGoods.Id, &miaoshaGoods.GoodsId, &miaoshaGoods.MiaoshaPrice, &miaoshaGoods.MiaoshaStock, &miaoshaGoods.StartTime, &miaoshaGoods.EndTime,
			&goods.Id, &goods.GoodsName, &goods.GoodsImg, &goods.GoodsPrice); err != nil {
			return
		}
		dto := model.MiaoshaGoodsDTO{
			MiaoshaGoods: &miaoshaGoods,
			Goods:        &goods,
		}
		goodsList = append(goodsList, &dto)
	}
	return
}

func (d *Dao) QueryById(goodsId int64) (goodsDTO *model.MiaoshaGoodsDTO, err error) {
	stmt, err := d.db.Prepare(_queryByIdSql)
	if err != nil {
		return
	}
	defer stmt.Close()
	miaoshaGoods := &model.MiaoshaGoods{}
	goods := &model.Goods{}
	err = stmt.QueryRow(goodsId).Scan(&miaoshaGoods.Id, &miaoshaGoods.GoodsId, &miaoshaGoods.MiaoshaPrice, &miaoshaGoods.MiaoshaStock, &miaoshaGoods.StartTime, &miaoshaGoods.EndTime,
		&goods.Id, &goods.GoodsName, &goods.GoodsImg, &goods.GoodsPrice)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	goodsDTO = &model.MiaoshaGoodsDTO{
		MiaoshaGoods: miaoshaGoods,
		Goods:        goods,
	}
	return
}
