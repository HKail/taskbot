package biz

import (
	"github.com/hkail/taskbot/app/biz/dao"
	"github.com/hkail/taskbot/app/botclient"
	"github.com/hkail/taskbot/app/conf"
)

type Biz struct {
	dao       *dao.Dao
	botClient *botclient.BotClient

	conf *conf.AppConf
}

func NewBiz(conf *conf.AppConf) (*Biz, error) {
	newDao, err := dao.NewDao(conf.DB)
	if err != nil {
		return nil, err
	}

	return &Biz{
		dao:  newDao,
		conf: conf,
	}, nil
}
