package models

import (
	"time"
)

type TbUser struct {
	Id       int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Username string    `json:"username" xorm:"not null comment('用户名') unique VARCHAR(50)"`
	Password string    `json:"password" xorm:"not null comment('密码，加密存储') VARCHAR(32)"`
	Phone    string    `json:"phone" xorm:"comment('注册手机号') unique VARCHAR(20)"`
	Created  time.Time `json:"created" xorm:"not null DATETIME"`
	Updated  time.Time `json:"updated" xorm:"not null DATETIME"`
}
