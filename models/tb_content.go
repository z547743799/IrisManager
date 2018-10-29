package models

import (
	"time"
)

type TbContent struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	CategoryId int64     `json:"category_id" xorm:"not null comment('内容类目ID') index BIGINT(20)"`
	Title      string    `json:"title" xorm:"comment('内容标题') VARCHAR(200)"`
	SubTitle   string    `json:"sub_title" xorm:"comment('子标题') VARCHAR(100)"`
	TitleDesc  string    `json:"title_desc" xorm:"comment('标题描述') VARCHAR(500)"`
	Url        string    `json:"url" xorm:"comment('链接') VARCHAR(500)"`
	Pic        string    `json:"pic" xorm:"comment('图片绝对路径') VARCHAR(300)"`
	Pic2       string    `json:"pic2" xorm:"comment('图片2') VARCHAR(300)"`
	Content    string    `json:"content" xorm:"comment('内容') TEXT"`
	Created    time.Time `json:"created" xorm:"DATETIME"`
	Updated    time.Time `json:"updated" xorm:"index DATETIME"`
}
