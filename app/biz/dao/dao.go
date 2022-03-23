package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hkail/taskbot/app/conf"
	"github.com/jmoiron/sqlx"
)

type Dao struct {
	db *sqlx.DB
}

func NewDao(dbConf *conf.DBConf) (*Dao, error) {
	db, err := sqlx.Connect("mysql", dbConf.MySQL.DataSource)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Dao{
		db: db,
	}, nil
}
