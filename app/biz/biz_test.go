package biz_test

import (
	"github.com/hkail/taskbot/app/biz"
	"github.com/hkail/taskbot/app/conf"
)

var defaultBiz *biz.Biz

func init() {
	newBiz, err := biz.NewBiz(&conf.AppConf{
		DB: &conf.DBConf{
			MySQL: &conf.MySQLConf{
				DataSource: "root:rootroot@tcp(127.0.0.1:3306)/taskbot?parseTime=true&charset=utf8mb4,utf8&autocommit=true&loc=Asia%2FShanghai",
			},
			Redis: nil,
		},
	})
	if err != nil {
		panic(err)
	}

	defaultBiz = newBiz
}

func getBiz() *biz.Biz {
	return defaultBiz
}
