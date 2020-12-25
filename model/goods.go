package model

import (
	"miaosha/conf"
	"time"
)

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

type MiaoshaGoodsDTO struct {
	Goods        *Goods
	MiaoshaGoods *MiaoshaGoods
}

type GoodsVO struct {
	Id         int64  `json:"id"`
	GoodsId    int64  `json:"goodsId"`
	Price      int64  `json:"price"`
	Stock      int    `json:"stock"`
	NowTime    int64  `json:"nowTime"`
	StartTime  int64  `json:"startTime"`
	EndTime    int64  `json:"endTime"`
	GoodsName  string `json:"goodsName"`
	GoodsImg   string `json:"goodsImg"`
	GoodsPrice int64  `json:"goodsPrice"`
	Status     int8   `json:"status"`
}

func (dto *MiaoshaGoodsDTO) ToGoodsVO() (goodsVo *GoodsVO) {
	goodsVo = &GoodsVO{}
	goodsVo.Id = dto.MiaoshaGoods.Id
	goodsVo.GoodsId = dto.MiaoshaGoods.GoodsId
	goodsVo.Price = dto.MiaoshaGoods.MiaoshaPrice
	goodsVo.Stock = dto.MiaoshaGoods.MiaoshaStock
	goodsVo.NowTime = time.Now().Unix()
	goodsVo.StartTime = dto.MiaoshaGoods.StartTime.Unix()
	goodsVo.EndTime = dto.MiaoshaGoods.EndTime.Unix()
	goodsVo.GoodsName = dto.Goods.GoodsName
	goodsVo.GoodsImg = dto.Goods.GoodsImg
	goodsVo.GoodsPrice = dto.Goods.GoodsPrice
	if goodsVo.NowTime < goodsVo.StartTime {
		goodsVo.Status = conf.MiaoshaStatusNotStarted
	} else if goodsVo.NowTime >= goodsVo.StartTime && goodsVo.NowTime <= goodsVo.EndTime {
		goodsVo.Status = conf.MiaoshaStatusOnGoing
	} else {
		goodsVo.Status = conf.MiaoshaStatusFinished
	}
	return
}
