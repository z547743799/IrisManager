package models

import (
	"time"
)

type TbItem struct {
	Id        int64     `json:"id" xorm:"pk comment('商品id，同时也是商品编号') BIGINT(20)"`
	Title     string    `json:"title" xorm:"not null comment('商品标题') VARCHAR(100)"`
	SellPoint string    `json:"sell_point" xorm:"comment('商品卖点') VARCHAR(500)"`
	Price     int64     `json:"price" xorm:"not null comment('商品价格，单位为：分') BIGINT(20)"`
	Num       int       `json:"num" xorm:"not null comment('库存数量') INT(10)"`
	Barcode   string    `json:"barcode" xorm:"comment('商品条形码') VARCHAR(30)"`
	Image     string    `json:"image" xorm:"comment('商品图片') VARCHAR(500)"`
	Cid       int64     `json:"cid" xorm:"not null comment('所属类目，叶子类目') index BIGINT(10)"`
	Status    int       `json:"status" xorm:"not null default 1 comment('商品状态，1-正常，2-下架，3-删除') index TINYINT(4)"`
	Created   time.Time `json:"created" xorm:"not null comment('创建时间') DATETIME"`
	Updated   time.Time `json:"updated" xorm:"not null comment('更新时间') index DATETIME"`
}
