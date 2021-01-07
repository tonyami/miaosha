package model

import (
	"miaosha/conf"
	"time"
)

// vo
type GoodsVO struct {
	Id          int64  `json:"id"`
	GoodsId     int64  `json:"goodsId"`
	Price       int64  `json:"price"`
	Duration    int64  `json:"duration"`
	GoodsName   string `json:"goodsName"`
	GoodsImg    string `json:"goodsImg"`
	OriginPrice int64  `json:"originPrice"`
	Status      int8   `json:"status"`
}

func (dto *GoodsDTO) ToVO() (vo *GoodsVO) {
	vo = &GoodsVO{}
	vo.Id = dto.MiaoshaGoods.Id
	vo.GoodsId = dto.MiaoshaGoods.GoodsId
	vo.Price = dto.MiaoshaGoods.MiaoshaPrice
	vo.GoodsName = dto.Goods.GoodsName
	vo.GoodsImg = dto.Goods.GoodsImg
	vo.OriginPrice = dto.Goods.GoodsPrice
	startTime := dto.MiaoshaGoods.StartTime.Unix()
	endTime := dto.MiaoshaGoods.EndTime.Unix()
	now := time.Now().Unix()
	if now < startTime {
		vo.Status = conf.StatusNotStarted
		vo.Duration = startTime - now
	} else if now >= startTime && now <= endTime {
		if dto.MiaoshaGoods.MiaoshaStock > 0 {
			vo.Status = conf.StatusOnGoing
		} else {
			vo.Status = conf.StatusSoldOut
		}
	} else {
		vo.Status = conf.StatusFinished
	}
	return
}

// dto
type GoodsDTO struct {
	Goods        *Goods
	MiaoshaGoods *MiaoshaGoods
}

// data model
type Goods struct {
	Id         int64
	GoodsName  string
	GoodsImg   string
	GoodsPrice int64
}

type MiaoshaGoods struct {
	Id           int64
	GoodsId      int64
	MiaoshaPrice int64
	MiaoshaStock int
	StartTime    time.Time
	EndTime      time.Time
}
