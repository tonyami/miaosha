package repository

import (
	"database/sql"
	"log"
	"miaosha/infra/db"
	"miaosha/model"
)

const (
	PageSize = 10
)

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
