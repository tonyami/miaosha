package service

type GoodsStatus int8 // 商品状态

const (
	Ended      GoodsStatus = -1 // 已结束
	NotStarted GoodsStatus = 0  // 未开始
	OnGoing    GoodsStatus = 1  // 进行中
	SoldOut    GoodsStatus = 2  // 已售罄
)
