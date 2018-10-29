package models

import (
	"time"
)

type TbOrder struct {
	OrderId      string    `json:"order_id" xorm:"not null pk default '' comment('订单id') VARCHAR(50)"`
	Payment      string    `json:"payment" xorm:"comment('实付金额。精确到2位小数;单位:元。如:200.07，表示:200元7分') VARCHAR(50)"`
	PaymentType  int       `json:"payment_type" xorm:"comment('支付类型，1、在线支付，2、货到付款') index INT(2)"`
	PostFee      string    `json:"post_fee" xorm:"comment('邮费。精确到2位小数;单位:元。如:200.07，表示:200元7分') VARCHAR(50)"`
	Status       int       `json:"status" xorm:"comment('状态：1、未付款，2、已付款，3、未发货，4、已发货，5、交易成功，6、交易关闭') index INT(10)"`
	CreateTime   time.Time `json:"create_time" xorm:"comment('订单创建时间') index DATETIME"`
	UpdateTime   time.Time `json:"update_time" xorm:"comment('订单更新时间') DATETIME"`
	PaymentTime  time.Time `json:"payment_time" xorm:"comment('付款时间') DATETIME"`
	ConsignTime  time.Time `json:"consign_time" xorm:"comment('发货时间') DATETIME"`
	EndTime      time.Time `json:"end_time" xorm:"comment('交易完成时间') DATETIME"`
	CloseTime    time.Time `json:"close_time" xorm:"comment('交易关闭时间') DATETIME"`
	ShippingName string    `json:"shipping_name" xorm:"comment('物流名称') VARCHAR(20)"`
	ShippingCode string    `json:"shipping_code" xorm:"comment('物流单号') VARCHAR(20)"`
	UserId       int64     `json:"user_id" xorm:"comment('用户id') BIGINT(20)"`
	BuyerMessage string    `json:"buyer_message" xorm:"comment('买家留言') VARCHAR(100)"`
	BuyerNick    string    `json:"buyer_nick" xorm:"comment('买家昵称') index VARCHAR(50)"`
	BuyerRate    int       `json:"buyer_rate" xorm:"comment('买家是否已经评价') INT(2)"`
}
