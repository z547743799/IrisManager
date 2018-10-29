package models

type TbOrderItem struct {
	Id       string `json:"id" xorm:"not null pk VARCHAR(20)"`
	ItemId   string `json:"item_id" xorm:"not null comment('商品id') index VARCHAR(50)"`
	OrderId  string `json:"order_id" xorm:"not null comment('订单id') index VARCHAR(50)"`
	Num      int    `json:"num" xorm:"comment('商品购买数量') INT(10)"`
	Title    string `json:"title" xorm:"comment('商品标题') VARCHAR(200)"`
	Price    int64  `json:"price" xorm:"comment('商品单价') BIGINT(50)"`
	TotalFee int64  `json:"total_fee" xorm:"comment('商品总金额') BIGINT(50)"`
	PicPath  string `json:"pic_path" xorm:"comment('商品图片地址') VARCHAR(200)"`
}
