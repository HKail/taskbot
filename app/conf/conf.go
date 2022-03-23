package conf

type AppConf struct {
	Bot *BotConf
	DB  *DBConf
}

type BotConf struct {
	AppID  string
	Token  string
	Secret string
}

type DBConf struct {
	MySQL *MySQLConf
	Redis *RedisConf
}

type MySQLConf struct {
	DataSource string
}

type RedisConf struct {
	Addr     string
	Password string
	DB       int
}
