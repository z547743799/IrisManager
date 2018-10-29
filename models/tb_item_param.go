package models

import (
	"time"
)

type TbItemParam struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ItemCatId int64     `json:"item_cat_id" xorm:"comment('商品类目ID') index BIGINT(20)"`
	ParamData string    `json:"param_data" xorm:"comment('参数数据，格式为json格式') TEXT"`
	Created   time.Time `json:"created" xorm:"DATETIME"`
	Updated   time.Time `json:"updated" xorm:"DATETIME"`
}
