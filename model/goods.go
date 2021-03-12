package model

import (
	"errors"
	"time"
)

const (
	Ended      int8 = -1 // 已结束
	NotStarted int8 = 0  // 未开始
	OnGoing    int8 = 1  // 进行中
	SoldOut    int8 = 2  // 已售罄
)

type Goods struct {
	Id          int64     `db:"id"`
	Name        string    `db:"name"`
	Img         string    `db:"img"`
	OriginPrice int64     `db:"origin_price"`
	Price       int64     `db:"price"`
	Stock       int       `db:"stock"`
	StartTime   time.Time `db:"start_time"`
	EndTime     time.Time `db:"end_time"`
}

type GoodsVO struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Img         string `json:"img"`
	OriginPrice int64  `json:"originPrice"`
	Price       int64  `json:"price"`
	Duration    int64  `json:"duration"`
	Status      int8   `json:"status"`
}

func (goods Goods) ToVO() GoodsVO {
	g := GoodsVO{}
	g.Id = goods.Id
	g.OriginPrice = goods.OriginPrice
	g.Name = goods.Name
	g.Img = goods.Img
	g.Price = goods.Price
	startTime := goods.StartTime.Unix()
	endTime := goods.EndTime.Unix()
	now := time.Now().Unix()
	if now < startTime {
		g.Status = NotStarted
		g.Duration = startTime - now
	} else if now >= startTime && now <= endTime {
		if goods.Stock > 0 {
			g.Status = OnGoing
		} else {
			g.Status = SoldOut
		}
	} else {
		g.Status = Ended
	}
	return g
}

func (goods Goods) Check() (err error) {
	now := time.Now().Unix()
	startTime := goods.StartTime.Unix()
	endTime := goods.EndTime.Unix()
	if now < startTime {
		err = errors.New("活动还未开始")
	} else if now > endTime {
		err = errors.New("活动已结束")
	} else if goods.Stock <= 0 {
		err = errors.New("商品已售罄")
	}
	return
}
