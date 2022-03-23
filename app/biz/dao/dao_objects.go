package dao

import "time"

// GuildUser guild_users 表结构体
type GuildUser struct {
	ID             int64     `db:"id"`
	GuildID        int64     `db:"guild_id"`
	UserID         int64     `db:"user_id"`
	ContCheckinCnt int       `db:"cont_checkin_cnt"`
	Experience     int       `db:"experience"`
	CreateTime     time.Time `db:"create_time"`
	UpdateTime     time.Time `db:"update_time"`
}

// GuildUserCheckin guild_users_checkin 表结构体
type GuildUserCheckin struct {
	ID         int64     `db:"id"`
	GuildID    int64     `db:"guild_id"`
	UserID     int64     `db:"user_id"`
	Yearmonth  int       `db:"yearmonth"`
	Days       uint32    `db:"days"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}
