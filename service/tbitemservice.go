package service

import (
	"IrisManager/db"
	"IrisManager/models"
	"github.com/go-xorm/xorm"
)

type TbItemService interface {
	GetAll() []models.TbItem
	Search(country string) []models.TbItem
	Get(id int64) *models.TbItem
	Delete(id int64) error
	Update(user *models.TbItem, columns []string) error
	Create(user *models.TbItem) error
}



type tbItemService struct {
	engine *xorm.Engine
}

func NewTbItemService()TbItemService {
	return &tbItemService{
		engine:db.X,
	}
}

func (d *tbItemService) Get(id int64) *models.TbItem {
	data := &models.TbItem{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *tbItemService) GetAll() []models.TbItem {
	datalist := make([]models.TbItem, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *tbItemService) Search(country string) []models.TbItem {
	datalist := make([]models.TbItem, 0)
	err := d.engine.Where("country=?", country).
		Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *tbItemService) Delete(id int64) error {
	data := &models.TbItem{Id:id}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *tbItemService) Update(data *models.TbItem, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *tbItemService) Create(data *models.TbItem) error {
	_, err := d.engine.Insert(data)
	return err
}
