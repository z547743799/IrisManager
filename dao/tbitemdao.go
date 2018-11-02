package dao

import (
	"github.com/go-xorm/xorm"
	"IrisManager/models"
	"IrisManager/db"
)

type TbItemDao struct {
	engine *xorm.Engine
}

func NewSuperstarDao( ) *TbItemDao {
	return &TbItemDao{
		engine:db.X,
	}
}

func (d *TbItemDao) Get(id int64) *models.TbItem {
	data := &models.TbItem{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *TbItemDao) GetAll() []models.TbItem {
	datalist := make([]models.TbItem, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *TbItemDao) Search(country string) []models.TbItem {
	datalist := make([]models.TbItem, 0)
	err := d.engine.Where("country=?", country).
		Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *TbItemDao) Delete(id int64) error {
	data := &models.TbItem{Id:id}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *TbItemDao) Update(data *models.TbItem, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *TbItemDao) Create(data *models.TbItem) error {
	_, err := d.engine.Insert(data)
	return err
}
