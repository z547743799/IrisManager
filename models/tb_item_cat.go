package models

import (
	"time"
)

type TbItemCat struct {
	Id        int64     `json:"id" xorm:"pk autoincr comment('类目ID') BIGINT(20)"`
	ParentId  int64     `json:"parent_id" xorm:"comment('父类目ID=0时，代表的是一级的类目') index(parent_id) BIGINT(20)"`
	Name      string    `json:"name" xorm:"comment('类目名称') VARCHAR(50)"`
	Status    int       `json:"status" xorm:"default 1 comment('状态。可选值:1(正常),2(删除)') index(parent_id) INT(1)"`
	SortOrder int       `json:"sort_order" xorm:"comment('排列序号，表示同级类目的展现次序，如数值相等则按名称次序排列。取值范围:大于零的整数') index INT(4)"`
	IsParent  int       `json:"is_parent" xorm:"default 1 comment('该类目是否为父类目，1为true，0为false') TINYINT(1)"`
	Created   time.Time `json:"created" xorm:"comment('创建时间') DATETIME"`
	Updated   time.Time `json:"updated" xorm:"comment('创建时间') DATETIME"`
}
