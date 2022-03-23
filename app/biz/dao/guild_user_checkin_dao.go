package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

// GetUserCheckinInfo 获取用户签到信息
// 当获取不到数据时返回 nil, nil
func (dao *Dao) GetUserCheckinInfo(ctx context.Context, t time.Time, gid, uid int) (*GuildUserCheckin, error) {
	yearMonth := getYearmonth(t)
	userCheckin := &GuildUserCheckin{}
	err := dao.db.GetContext(ctx, userCheckin, "SELECT * FROM guild_users_checkin WHERE guild_id = ? AND user_id = ? AND yearmonth = ? LIMIT 1",
		gid, uid, yearMonth)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		return nil, nil
	}

	return userCheckin, nil
}

// Checkin 用户签到
func (dao *Dao) Checkin(ctx context.Context, t time.Time, gid, uid int) (todayFirstCheckin bool, err error) {
	yearMonth := getYearmonth(t)
	dayMask := getDayMask(t)
	todayFirstCheckin = true

	err = dao.txx(ctx, func(tx *sqlx.Tx) error {
		checkinInfo, err := dao.txGetUserCheckinForUpdate(ctx, tx, yearMonth, gid, uid)
		if err != nil {
			return err
		}

		if checkinInfo == nil { // 本月第一次签到
			return dao.txCreateUserCheckinInfo(ctx, tx, yearMonth, dayMask, gid, uid)
		}

		if checkinInfo.Days&dayMask > 0 { // 当日已签到
			todayFirstCheckin = false
			return nil
		}

		return dao.txUpdateUserCheckinInfo(ctx, tx, yearMonth, dayMask, gid, uid)
	})
	if err != nil {
		return
	}

	return
}

// ListUserCheckinInfos 获取用户历史签到信息
func (dao *Dao) ListUserCheckinInfos(ctx context.Context, gid, uid int) ([]GuildUserCheckin, error) {
	// 对于一个频道而言, 一个用户一年最多也只会有 12 条记录, 因此无需考虑分批获取
	rows, err := dao.db.QueryxContext(ctx, "SELECT * FROM guild_users_checkin WHERE guild_id = ? AND user_id = ?", gid, uid)
	if err != nil {
		return nil, err
	}

	userCheckins := make([]GuildUserCheckin, 0)
	for rows.Next() {
		userCheckin := GuildUserCheckin{}
		err = rows.StructScan(&userCheckin)
		if err != nil {
			return nil, err
		}

		userCheckins = append(userCheckins, userCheckin)
	}

	return userCheckins, nil
}

func (dao *Dao) txGetUserCheckinForUpdate(ctx context.Context, tx sqlx.QueryerContext, yearMonth int, gid, uid int) (*GuildUserCheckin, error) {
	userCheckin := &GuildUserCheckin{}
	err := sqlx.GetContext(ctx, tx, userCheckin, "SELECT * FROM guild_users_checkin WHERE guild_id = ? AND user_id = ? AND yearmonth = ? LIMIT 1",
		gid, uid, yearMonth)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		return nil, nil
	}

	return userCheckin, nil
}

func (dao *Dao) txCreateUserCheckinInfo(ctx context.Context, tx sqlx.ExecerContext, yearMonth int, dayMask uint32, gid, uid int) error {
	query := "INSERT INTO guild_users_checkin(guild_id, user_id, yearmonth, days) VALUES(?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, query, gid, uid, yearMonth, dayMask)

	return err
}

func (dao *Dao) txUpdateUserCheckinInfo(ctx context.Context, tx sqlx.ExecerContext, yearMonth int, dayMask uint32, gid, uid int) error {
	query := "UPDATE guild_users_checkin SET days=(days|?) WHERE guild_id = ? AND user_id = ? AND yearmonth = ?"
	_, err := tx.ExecContext(ctx, query, dayMask, gid, uid, yearMonth)

	return err
}

func (dao *Dao) txx(ctx context.Context, fn func(tx *sqlx.Tx) error) error {
	tx, err := dao.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	// commit 和 rollback 只有一个会成功执行
	defer tx.Rollback()

	err = fn(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func getYearmonth(t time.Time) int {
	return t.Year()*100 + int(t.Month())
}

// getDayMask 获取当天日期 bitmap 掩码
func getDayMask(t time.Time) uint32 {
	return 1 << (t.Day() - 1)
}
