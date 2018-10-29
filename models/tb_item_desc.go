package models

import (
	"time"
)

type TbItemDesc struct {
	ItemId   int64     `json:"item_id" xorm:"not null pk comment('商品ID') BIGINT(20)"`
	ItemDesc string    `json:"item_desc" xorm:"comment('商品描述') TEXT"`
	Created  time.Time `json:"created" xorm:"comment('创建时间') DATETIME"`
	Updated  time.Time `json:"updated" xorm:"comment('更新时间') DATETIME"`
}
