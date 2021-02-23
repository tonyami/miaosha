package service

import (
	"errors"
)

var IGoodsService GoodsService

func GetGoodsService() GoodsService {
	return IGoodsService
}

type GoodsService interface {
	GetList(int) ([]*GoodsDTO, error)
	Get(int64) (*GoodsDTO, error)
}

type GoodsDTO struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	Img         string      `json:"img"`
	OriginPrice int64       `json:"originPrice"`
	Price       int64       `json:"price"`
	Duration    int64       `json:"duration"`
	Status      GoodsStatus `json:"status"`
}

func (dto *GoodsDTO) Check() (err error) {
	if dto.Status == NotStarted {
		err = errors.New("秒杀还未开始")
	} else if dto.Status == SoldOut {
		err = errors.New("商品已售罄")
	} else if dto.Status == Finished {
		err = errors.New("秒杀已结束")
	}
	return
}
