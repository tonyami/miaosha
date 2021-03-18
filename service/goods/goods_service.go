package goods

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/infra/cache"
	"miaosha/infra/code"
	"miaosha/model"
	"miaosha/repository"
	"miaosha/service"
	"sync"
)

var once sync.Once

func InitService() {
	once.Do(func() {
		service.GoodsService = &goodsService{
			goodsRepository: repository.NewGoodsRepository(),
			redis:           cache.Client,
		}
	})
}

type goodsService struct {
	goodsRepository repository.GoodsRepository
	redis           *redis.Client
}

func (s *goodsService) InitGoodsStock() (err error) {
	var list []model.Goods
	if list, err = s.goodsRepository.GetGoodsList(-1); err != nil {
		err = code.DBErr
		return
	}
	for _, v := range list {
		if err = s.SetGoodsStock(v.Id, v.Stock); err != nil {
			return
		}
	}
	return
}

func (s *goodsService) GetGoodsList(page int) (list []model.GoodsVO, err error) {
	list = make([]model.GoodsVO, 0)
	var goodsList []model.Goods
	if goodsList, err = s.goodsRepository.GetGoodsList(page); err != nil {
		err = code.DBErr
		return
	}
	for _, v := range goodsList {
		list = append(list, v.ToVO())
	}
	return
}

func (s *goodsService) GetGoodsVO(id int64) (goodsVO model.GoodsVO, err error) {
	var goods model.Goods
	if goods, err = s.GetGoods(id); err != nil {
		err = code.DBErr
	}
	goodsVO = goods.ToVO()
	return
}

func (s *goodsService) GetGoods(id int64) (goods model.Goods, err error) {
	if goods, err = s.goodsRepository.GetGoods(id); err != nil {
		err = code.DBErr
	}
	return
}

func (s *goodsService) SetGoodsStock(goodsId int64, stock int) (err error) {
	if err = s.redis.Set(service.Ctx, fmt.Sprintf(service.GoodsStockKey, goodsId), stock, -1).Err(); err != nil {
		log.Printf("redis.Set() failed, err: %v", err)
		err = code.RedisErr
	}
	return
}

func (s *goodsService) DecrStock(goodsId int64) (stock int64, err error) {
	if stock, err = s.redis.Decr(service.Ctx, fmt.Sprintf(service.GoodsStockKey, goodsId)).Result(); err != nil {
		log.Printf("redis.Decr() failed, err: %v", err)
		err = code.RedisErr
	}
	return
}

func (s *goodsService) IncrStock(goodsId int64) (err error) {
	if err = s.redis.Incr(service.Ctx, fmt.Sprintf(service.GoodsStockKey, goodsId)).Err(); err != nil {
		log.Printf("redis.Incr() failed, err: %v", err)
		err = code.RedisErr
	}
	return
}
