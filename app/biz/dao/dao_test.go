package dao_test

import (
	"github.com/hkail/taskbot/app/biz/dao"
	"github.com/hkail/taskbot/app/conf"
)

var defaultDao *dao.Dao

func init() {
	newDao, err := dao.NewDao(&conf.DBConf{
		MySQL: &conf.MySQLConf{
			DataSource: "root:rootroot@tcp(127.0.0.1:3306)/taskbot?parseTime=true&charset=utf8mb4,utf8&autocommit=true&loc=Asia%2FShanghai",
		},
		Redis: nil,
	})
	if err != nil {
		panic(err)
	}

	defaultDao = newDao
}

func getDao() *dao.Dao {
	return defaultDao
}
