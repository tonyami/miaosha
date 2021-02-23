package goods

import (
	"errors"
	"log"
	"miaosha/internal/db"
	"miaosha/service"
	"sync"
)

var once sync.Once

func init() {
	log.Printf("init goods service...")
	once.Do(func() {
		service.IGoodsService = new(goodsService)
	})
}

type goodsService struct {
}

func (*goodsService) GetList(page int) (list []*service.GoodsDTO, err error) {
	dao := NewDao(db.Get())
	var rows []*Goods
	if rows, err = dao.GetList(page, service.PageSize); err != nil {
		err = errors.New("db error")
		return
	}
	list = []*service.GoodsDTO{}
	for _, row := range rows {
		list = append(list, row.toDTO())
	}
	return
}

func (*goodsService) Get(goodsId int64) (goods *service.GoodsDTO, err error) {
	dao := NewDao(db.Get())
	var row *Goods
	if row, err = dao.Get(goodsId); err != nil {
		err = errors.New("db error")
		return
	}
	goods = row.toDTO()
	return
}
