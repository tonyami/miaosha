package goods

import (
	"log"
	"miaosha/conf"
	"miaosha/conf/errmsg"
	"miaosha/dao/goods"
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
	goodsDTOList, err := s.dao.QueryAll(page, conf.GoodsSize)
	if err != nil {
		log.Printf("【商品】查询列表失败: %s, %d\n", err, page)
		err = errmsg.SystemErr
		return
	}
	goodsList = []*model.GoodsVO{}
	for _, goodsDTO := range goodsDTOList {
		goodsList = append(goodsList, goodsDTO.ToGoodsVO())
	}
	return
}

func (s *Service) GetGoodsDetail(goodsId int64) (goods *model.GoodsVO, err error) {
	goodsDTO, err := s.dao.QueryById(goodsId)
	if err != nil {
		log.Printf("【商品】查询详情失败: %s, %d\n", err, goodsId)
		err = errmsg.SystemErr
		return
	}
	if goodsDTO == nil {
		err = errmsg.InvalidParameter
		return
	}
	goods = goodsDTO.ToGoodsVO()
	return
}
