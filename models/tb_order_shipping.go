package models

import (
	"time"
)

type TbOrderShipping struct {
	OrderId          string    `json:"order_id" xorm:"not null pk comment('订单ID') VARCHAR(50)"`
	ReceiverName     string    `json:"receiver_name" xorm:"comment('收货人全名') VARCHAR(20)"`
	ReceiverPhone    string    `json:"receiver_phone" xorm:"comment('固定电话') VARCHAR(20)"`
	ReceiverMobile   string    `json:"receiver_mobile" xorm:"comment('移动电话') VARCHAR(30)"`
	ReceiverState    string    `json:"receiver_state" xorm:"comment('省份') VARCHAR(10)"`
	ReceiverCity     string    `json:"receiver_city" xorm:"comment('城市') VARCHAR(10)"`
	ReceiverDistrict string    `json:"receiver_district" xorm:"comment('区/县') VARCHAR(20)"`
	ReceiverAddress  string    `json:"receiver_address" xorm:"comment('收货地址，如：xx路xx号') VARCHAR(200)"`
	ReceiverZip      string    `json:"receiver_zip" xorm:"comment('邮政编码,如：310001') VARCHAR(6)"`
	Created          time.Time `json:"created" xorm:"DATETIME"`
	Updated          time.Time `json:"updated" xorm:"DATETIME"`
}
