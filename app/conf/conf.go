package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type AppConf struct {
	Bot *BotConf `yaml:"bot"`
	DB  *DBConf  `yaml:"db"`
}

type BotConf struct {
	IsSandbox bool   `yaml:"is_sandbox"`
	AppID     uint64 `yaml:"app_id"`
	Token     string `yaml:"token"`
	Secret    string `yaml:"secret"`
}

type DBConf struct {
	MySQL *MySQLConf `yaml:"mysql"`
	Redis *RedisConf `yaml:"redis"`
}

type MySQLConf struct {
	DataSource string `yaml:"data_source"`
}

type RedisConf struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// LoadConfig 加载配置
func LoadConfig(confPath string) (*AppConf, error) {
	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, err
	}

	appConf := AppConf{}
	err = yaml.Unmarshal(data, &appConf)
	if err != nil {
		return nil, err
	}

	return &appConf, nil
}
