package model

import (
	"miaosha/conf"
	"miaosha/internal/code"
	"time"
)

// vo
type GoodsVO struct {
	Id          int64              `json:"id"`
	Name        string             `json:"name"`
	Img         string             `json:"img"`
	OriginPrice int64              `json:"originPrice"`
	Price       int64              `json:"price"`
	Duration    int64              `json:"duration"`
	Status      conf.MiaoshaStatus `json:"status"`
}

func (goods *Goods) ToVO() (goodsVO *GoodsVO) {
	goodsVO = &GoodsVO{}
	goodsVO.Id = goods.Id
	goodsVO.OriginPrice = goods.OriginPrice
	goodsVO.Name = goods.Name
	goodsVO.Img = goods.Img
	goodsVO.Price = goods.Price
	startTime := goods.StartTime.Unix()
	endTime := goods.EndTime.Unix()
	now := time.Now().Unix()
	if now < startTime {
		goodsVO.Status = conf.MiaoshaNotStarted
		goodsVO.Duration = startTime - now
	} else if now >= startTime && now <= endTime {
		if goods.Stock > 0 {
			goodsVO.Status = conf.MiaoshaOnGoing
		} else {
			goodsVO.Status = conf.MiaoshaSoldOut
		}
	} else {
		goodsVO.Status = conf.MiaoshaFinished
	}
	return
}

func (goods *Goods) Check() (err error) {
	now := time.Now().Unix()
	startTime := goods.StartTime.Unix()
	endTime := goods.EndTime.Unix()
	if now < startTime {
		err = code.MiaoshaNotStarted
	} else if now > endTime {
		err = code.MiaoshaFinished
	} else if goods.Stock <= 0 {
		err = code.MiaoshaSoldOut
	}
	return
}

// data model
type Goods struct {
	Id          int64
	Name        string
	Img         string
	OriginPrice int64
	Price       int64
	Stock       int
	StartTime   time.Time
	EndTime     time.Time
}
