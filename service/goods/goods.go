package goods

import (
	"log"
	"miaosha/conf"
	"miaosha/dao/goods"
	"miaosha/internal/code"
	"miaosha/model"
)

type Service struct {
	dao *goods.Dao
}

func New() *Service {
	return &Service{
		dao: goods.New(),
	}
}

func (s *Service) GetList(page int) (goodsList []*model.GoodsVO, err error) {
	goodsDTOList, err := s.dao.GetList(page, conf.PageSize)
	if err != nil {
		log.Printf("【Goods】GetList Failed: %s", err)
		err = code.SystemErr
		return
	}
	goodsList = []*model.GoodsVO{}
	for _, goodsDTO := range goodsDTOList {
		goodsList = append(goodsList, goodsDTO.ToVO())
	}
	return
}

func (s *Service) Get(goodsId int64) (goods *model.Goods, err error) {
	if goods, err = s.dao.Get(goodsId); err != nil {
		log.Printf("【Goods】Get Failed: %s", err)
		err = code.SystemErr
		return
	}
	if goods == nil {
		err = code.GoodsNotFound
		return
	}
	return
}
