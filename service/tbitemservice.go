package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-xorm/xorm"
	"gitlab.com/z547743799/iriscommon/pojo"
	"gitlab.com/z547743799/iriscontent/redisinit"
	"gitlab.com/z547743799/irismanager/db"
	"gitlab.com/z547743799/irismanager/models"
)

type TbItemService interface {
	Get(id int64) *models.TbItem
	GetList(page int, rows int) *pojo.EasyUIDataGridResult
	GetItemDescById(itemid int64) *models.TbItemDesc
}

type tbItemService struct {
	engine *xorm.Engine
}

func NewTbItemService() TbItemService {
	return &tbItemService{
		engine: db.X,
	}
}

func (d *tbItemService) Get(id int64) *models.TbItem {
	Data := &models.TbItem{}

	pool := redisinit.Re.Get()

	item, err := pool.Do("get", fmt.Sprintln(id))

	if item != nil || err == nil {
		err = json.Unmarshal([]byte(item.(string)), Data)
		if err == nil {
			return Data
		}
	}

	ok, err := d.engine.Where("id=?", id).Get(Data)

	if ok && err == nil || Data != nil {

		var data, err = json.Marshal(Data)
		if err != nil {
			return nil
		}

		pool.Do("set", fmt.Sprintln(id), string(data), "EX", 30*time.Second)

		return Data
	} else {
		Data.Id = 0
		return Data
	}
}
func (d *tbItemService) GetList(page int, rows int) *pojo.EasyUIDataGridResult {
	datalist := make([]models.TbItem, 0)
	err := d.engine.Desc("id").Find(&datalist)

	leng := len(datalist)

	result := new(pojo.EasyUIDataGridResult)

	result.Total = leng
	result.Rows = datalist[(page-1)*rows : page*rows]
	if err != nil {
		return result
	} else {
		return result
	}

}
func (d *tbItemService) GetItemDescById(itemid int64) *models.TbItemDesc {
	ItemDesc := &models.TbItemDesc{}
	pool := redisinit.Re.Get()

	item, err := pool.Do("get", fmt.Sprintln(itemid))

	if item != nil || err == nil {
		err = json.Unmarshal([]byte(item.(string)), ItemDesc)
		if err == nil {
			return ItemDesc
		}
	}
	ok, err := d.engine.Where("item_id=?", itemid).Get(ItemDesc)

	if ok && err == nil || ItemDesc != nil {

		var data, err = json.Marshal(ItemDesc)
		if err != nil {
			return nil
		}

		pool.Do("set", fmt.Sprintln(itemid), string(data), "EX", 30*time.Second)

		return ItemDesc
	} else {
		ItemDesc.ItemId = 0
		return ItemDesc
	}
}
