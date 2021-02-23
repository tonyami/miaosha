package goods

import (
	"database/sql"
	"log"
)

type Dao struct {
	db *sql.DB
}

func NewDao(db *sql.DB) *Dao {
	return &Dao{
		db: db,
	}
}

var (
	_getListSql   = "select `id`, `name`, `img`, `origin_price`, `price`, `stock`, `start_time`, `end_time` from `miaosha_goods` order by `id` desc limit ?, ?"
	_getSql       = "select `id`, `name`, `img`, `origin_price`, `price`, `stock`, `start_time`, `end_time` from `miaosha_goods` where `id` = ? limit 1"
	_decrStockSql = "update `miaosha_goods` set `stock` = `stock` - 1 where `id` = ? and `stock` > 0"
	_incrStockSql = "update `miaosha_goods` set `stock` = `stock` + 1 where `id` = ?"
)

func (dao *Dao) GetList(page, size int) (list []*Goods, err error) {
	var rows *sql.Rows
	if rows, err = dao.db.Query(_getListSql, (page-1)*size, size); err != nil {
		log.Printf("dao.db.Query(_getListSql, %d, %d) failed, err: %v", (page-1)*size, size, err)
		return
	}
	defer rows.Close()
	list = []*Goods{}
	for rows.Next() {
		goods := new(Goods)
		if err = rows.Scan(&goods.Id, &goods.Name, &goods.Img, &goods.OriginPrice, &goods.Price, &goods.Stock, &goods.StartTime, &goods.EndTime); err != nil {
			log.Printf("rows.Scan() failed, err: %v", err)
			return
		}
		list = append(list, goods)
	}
	return
}

func (dao *Dao) Get(goodsId int64) (goods *Goods, err error) {
	var stmt *sql.Stmt
	if stmt, err = dao.db.Prepare(_getSql); err != nil {
		log.Printf("dao.db.Prepare(%d) failed, err: %v", goodsId, err)
		return
	}
	defer stmt.Close()
	goods = &Goods{}
	if err = stmt.QueryRow(goodsId).Scan(&goods.Id, &goods.Name, &goods.Img, &goods.OriginPrice, &goods.Price, &goods.Stock, &goods.StartTime, &goods.EndTime); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Printf("stmt.QueryRow(%d) failed, err: %v", goodsId, err)
		}
	}
	return
}

func (dao *Dao) DecrStock(goodsId int64) (err error) {
	if _, err = dao.db.Exec(_decrStockSql, goodsId); err != nil {
		log.Printf("dao.db.Exec(_decrStockSql, %d) failed, err: %v", goodsId, err)
		return
	}
	return
}

func (dao *Dao) IncrStock(goodsId int64) (err error) {
	if _, err = dao.db.Exec(_incrStockSql, goodsId); err != nil {
		log.Printf("dao.db.Exec(_incrStockSql, %d) failed, err: %v", goodsId, err)
		return
	}
	return
}
