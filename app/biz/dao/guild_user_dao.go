package dao

import (
	"context"
	"database/sql"
)

// CreateGuildUser 创建频道用户信息
func (dao *Dao) CreateGuildUser(ctx context.Context, gid, uid int64) (int64, error) {
	result, err := dao.db.ExecContext(ctx, "INSERT INTO guild_users (guild_id, user_id) VALUES (?, ?)", gid, uid)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// GetGuildUserByGidAndUid 根据频道id和用户id获取用户信息
// 当获取不到数据时返回 nil, nil
func (dao *Dao) GetGuildUserByGidAndUid(ctx context.Context, gid, uid int64) (*GuildUser, error) {
	user := &GuildUser{}
	err := dao.db.GetContext(ctx, user, "SELECT * FROM guild_users WHERE guild_id = ? AND user_id = ?", gid, uid)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		return nil, nil
	}

	return user, nil
}
