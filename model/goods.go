package model

import (
	"miaosha/conf"
	"time"
)

// vo
type GoodsVO struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Img         string `json:"img"`
	OriginPrice int64  `json:"originPrice"`
	Price       int64  `json:"price"`
	Duration    int64  `json:"duration"`
	Status      int8   `json:"status"`
}

func (goods *Goods) ToVO() (goodsVO *GoodsVO) {
	goodsVO = &GoodsVO{}
	goodsVO.Id = goods.Id
	goodsVO.OriginPrice = goods.OriginPrice
	goodsVO.Name = goods.Name
	goodsVO.Img = goods.Img
	goodsVO.OriginPrice = goods.OriginPrice
	startTime := goods.StartTime.Unix()
	endTime := goods.EndTime.Unix()
	now := time.Now().Unix()
	if now < startTime {
		goodsVO.Status = conf.StatusNotStarted
		goodsVO.Duration = startTime - now
	} else if now >= startTime && now <= endTime {
		if goods.Stock > 0 {
			goodsVO.Status = conf.StatusOnGoing
		} else {
			goodsVO.Status = conf.StatusSoldOut
		}
	} else {
		goodsVO.Status = conf.StatusFinished
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
