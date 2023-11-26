package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
	"miaosha/infra/db"
	"miaosha/model"
)

type GoodsRepository interface {
	GetGoodsList(page int) (list []model.Goods, err error)
	GetGoods(id int64) (goods model.Goods, err error)
}

type goodsRepository struct {
	db *sqlx.DB
}

func NewGoodsRepository() GoodsRepository {
	return &goodsRepository{db: db.DB}
}

func (r *goodsRepository) GetGoodsList(page int) (list []model.Goods, err error) {
	args := []interface{}{}
	s := ""
	if page > 0 {
		s = "limit ?, ?"
		args = append(args, (page-1)*10)
		args = append(args, 10)
	}
	sqlStr := "select * from miaosha_goods order by id desc "
	if err = r.db.Select(&list, sqlStr+s, args...); err != nil {
		log.Printf("db.Select() failed, err: %v", err)
	}
	return
}

func (r *goodsRepository) GetGoods(id int64) (goods model.Goods, err error) {
	if err = r.db.Get(&goods, "select * from miaosha_goods where id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Printf("db.Get() failed, err: %v", err)
		}
	}
	return
}
