package goods

import (
	"errors"
	"miaosha/service"
	"time"
)

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

func (goods *Goods) toDTO() *service.GoodsDTO {
	dto := new(service.GoodsDTO)
	dto.Id = goods.Id
	dto.OriginPrice = goods.OriginPrice
	dto.Name = goods.Name
	dto.Img = goods.Img
	dto.Price = goods.Price
	startTime := goods.StartTime.Unix()
	endTime := goods.EndTime.Unix()
	now := time.Now().Unix()
	if now < startTime {
		dto.Status = service.NotStarted
		dto.Duration = startTime - now
	} else if now >= startTime && now <= endTime {
		if goods.Stock > 0 {
			dto.Status = service.OnGoing
		} else {
			dto.Status = service.SoldOut
		}
	} else {
		dto.Status = service.Ended
	}
	return dto
}

func (goods *Goods) Check() (err error) {
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
