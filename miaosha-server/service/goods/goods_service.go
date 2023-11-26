package goods

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/infra/cache"
	"miaosha/infra/code"
	"miaosha/model"
	"miaosha/repository"
	"miaosha/service"
	"sync"
	"time"
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
	// 尝试从缓存中取
	if goods, err = s.getGoodsFromCache(id); err != nil {
		return
	}
	// 从缓存中取到了
	if goods.Id > 0 {
		return
	}
	// 从数据库中取
	if goods, err = s.goodsRepository.GetGoods(id); err != nil {
		err = code.DBErr
		return
	}
	// 再放入缓存
	err = s.setGoodsCache(goods)
	return
}

func (s *goodsService) setGoodsCache(goods model.Goods) (err error) {
	var data []byte
	if data, err = json.Marshal(goods); err != nil {
		log.Printf("json.Marshal() failed, err: %v", err)
		err = code.SerializeErr
		return
	}
	if err = s.redis.Set(service.Ctx, fmt.Sprintf(service.GoodsKey, goods.Id), string(data), 12*time.Hour).Err(); err != nil {
		log.Printf("redis.Set() failed, err: %v", err)
		err = code.RedisErr
	}
	return
}

func (s *goodsService) getGoodsFromCache(id int64) (goods model.Goods, err error) {
	var res string
	if res, err = s.redis.Get(service.Ctx, fmt.Sprintf(service.GoodsKey, id)).Result(); err != nil {
		if err == redis.Nil {
			err = nil
		} else {
			log.Printf("redis.Set() failed, err: %v", err)
			err = code.RedisErr
		}
		return
	}
	if err = json.Unmarshal([]byte(res), &goods); err != nil {
		log.Printf("json.Unmarshal() failed, err: %v, json: %#v", err, goods)
		err = code.SerializeErr
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
