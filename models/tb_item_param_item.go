package models

import (
	"time"
)

type TbItemParamItem struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ItemId    int64     `json:"item_id" xorm:"comment('商品ID') index BIGINT(20)"`
	ParamData string    `json:"param_data" xorm:"comment('参数数据，格式为json格式') TEXT"`
	Created   time.Time `json:"created" xorm:"DATETIME"`
	Updated   time.Time `json:"updated" xorm:"DATETIME"`
}
