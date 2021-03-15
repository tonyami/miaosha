package repository

import (
	"database/sql"
	"fmt"
	"log"
	"miaosha/infra/db"
	"miaosha/infra/rdb"
	"miaosha/model"
)

const (
	PageSize = 10
)

const goodsStockKey = "goods_stock:%d"

func SetGoodsStock(goodsId int64, stock int) (err error) {
	if err = rdb.Conn().Set(ctx, fmt.Sprintf(goodsStockKey, goodsId), stock, -1).Err(); err != nil {
		log.Printf("rdb.Set() failed, err: %v", err)
	}
	return
}

func DecrStock(goodsId int64) (stock int64, err error) {
	if stock, err = rdb.Conn().Decr(ctx, fmt.Sprintf(goodsStockKey, goodsId)).Result(); err != nil {
		log.Printf("rdb.Decr() failed, err: %v", err)
		return
	}
	return
}

func IncrStock(goodsId int64) (err error) {
	if err = rdb.Conn().Incr(ctx, fmt.Sprintf(goodsStockKey, goodsId)).Err(); err != nil {
		log.Printf("rdb.Incr() failed, err: %v", err)
		return
	}
	return
}

func GetGoodsList(page int) (list []model.Goods, err error) {
	if err = db.Conn().Select(&list, "select * from miaosha_goods order by id desc limit ?, ?", (page-1)*PageSize, PageSize); err != nil {
		log.Printf("dao.GetGoodsList() failed, err: %v", err)
	}
	return
}

func GetGoods(id int64) (goods model.Goods, err error) {
	if err = db.Conn().Get(&goods, "select * from miaosha_goods where id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Printf("dao.GetGoods() failed, err: %v", err)
		}
	}
	return
}
