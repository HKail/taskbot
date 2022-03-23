package biz

import (
	"context"
	"log"
	"time"

	"github.com/hkail/taskbot/app/biz/dao"
	"github.com/hkail/taskbot/app/util"
)

// UserCheckin 用户签到
func (biz *Biz) UserCheckin(ctx context.Context, t time.Time, gid, uid int) (todayFirstCheckin bool, err error) {
	// TODO 使用缓存判断用户当天是否已签到

	return biz.dao.Checkin(ctx, t, gid, uid)
}

// GetUserContCheckinDays 获取用户连续签到天数
func (biz *Biz) GetUserContCheckinDays(ctx context.Context, gid, uid int) (int, error) {
	userCheckins, err := biz.dao.ListUserCheckinInfos(ctx, gid, uid)
	if err != nil {
		log.Printf("GetUserContCheckinDays.ListUserCheckinInfos has err=%v", err)
		return 0, err
	}

	if len(userCheckins) == 0 {
		return 0, nil
	}

	checkinDateStatusMapper := GetUserCheckinDateStatusMapper(userCheckins)
	dateTime := time.Now()
	count := 0
	if checkinDateStatusMapper[util.GetDateNumberVal(dateTime)] { // 今日已签到
		count = 1
	}

	for {
		dateTime = dateTime.AddDate(0, 0, -1)
		if !checkinDateStatusMapper[util.GetDateNumberVal(dateTime)] {
			break
		}

		count++
	}

	return count, nil
}

// GetUserCheckinDateStatusMapper 获取用户签到日期信息
// 返回的数据格式如: {20060102: true, 20060103: false, ...}
func GetUserCheckinDateStatusMapper(userCheckins []dao.GuildUserCheckin) map[int]bool {
	dateStatusMapper := make(map[int]bool)
	for _, userCheckin := range userCheckins {
		numOfDays := util.GetDaysOfYearmonth(userCheckin.Yearmonth)
		highPosDateVal := userCheckin.Yearmonth * 100

		for i := 0; i < numOfDays; i++ {
			dateStatusMapper[highPosDateVal+i+1] = userCheckin.Days&(1<<i) > 0
		}
	}

	return dateStatusMapper
}
