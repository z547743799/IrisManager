package db

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// X 全局DB
var X *xorm.Engine

func init() {
	var err error
	cfg, err := ini.Load("/home/lzw/DarkShell/src/gitlab.com/z547743799/irismanager/config/db.ini")
	if err != nil {
		log.Fatal(err)
	}

	username := cfg.Section("mysql").Key("username").Value()
	password := cfg.Section("mysql").Key("password").Value()
	url := cfg.Section("mysql").Key("url").Value()

	source := fmt.Sprintf("%s:%s@%s", username, password, url)
	X, err = xorm.NewEngine("mysql", source)

	// 30minute ping db to keep connection
	timer := time.NewTicker(time.Minute * 30)
	go func(x *xorm.Engine) {
		for _ = range timer.C {
			err = x.Ping()
			if err != nil {
				log.Fatalf("db connect error: %#v\n", err.Error())
			}
		}
	}(X)
	// x.ShowSQL(true)
}
