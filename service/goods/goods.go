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

func (s *Service) GetGoodsList(page int) (goodsList []*model.GoodsVO, err error) {
	goodsDTOList, err := s.dao.GetList(page, conf.PageSize)
	if err != nil {
		log.Printf("GetGoodsList Failed: %s", err)
		err = code.SystemErr
		return
	}
	goodsList = []*model.GoodsVO{}
	for _, goodsDTO := range goodsDTOList {
		goodsList = append(goodsList, goodsDTO.ToVO())
	}
	return
}

func (s *Service) GetGoods(goodsId int64) (goods *model.GoodsVO, err error) {
	goodsDTO, err := s.dao.Get(goodsId)
	if err != nil {
		log.Printf("GetGoods Failed: %s", err)
		err = code.SystemErr
		return
	}
	if goodsDTO == nil {
		err = code.GoodsNotFound
		return
	}
	goods = goodsDTO.ToVO()
	return
}
