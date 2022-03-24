package dao_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDao_Checkin(t *testing.T) {
	dao := getDao()
	ctx := context.Background()

	type args struct {
		ctx context.Context
		t   time.Time
		gid uint64
		uid uint64
	}
	tests := []struct {
		name                  string
		args                  args
		wantTodayFirstCheckin bool
		wantErr               bool
	}{
		{
			name: "测试正常签到-1",
			args: args{
				ctx: ctx,
				t:   time.Date(2022, 3, 23, 0, 0, 0, 0, time.Local),
				gid: 1,
				uid: 1,
			},
			wantTodayFirstCheckin: true,
			wantErr:               false,
		},
		{
			name: "测试正常签到-2",
			args: args{
				ctx: ctx,
				t:   time.Date(2022, 3, 22, 0, 0, 0, 0, time.Local),
				gid: 1,
				uid: 1,
			},
			wantTodayFirstCheckin: true,
			wantErr:               false,
		},
		{
			name: "测试正常签到-3",
			args: args{
				ctx: ctx,
				t:   time.Date(2022, 3, 20, 0, 0, 0, 0, time.Local),
				gid: 1,
				uid: 1,
			},
			wantTodayFirstCheckin: true,
			wantErr:               false,
		},
		{
			name: "测试非正常签到-1",
			args: args{
				ctx: ctx,
				t:   time.Date(2022, 3, 23, 0, 0, 0, 0, time.Local),
				gid: 1,
				uid: 1,
			},
			wantTodayFirstCheckin: false,
			wantErr:               false,
		},
		{
			name: "测试非正常签到-2",
			args: args{
				ctx: ctx,
				t:   time.Date(2022, 3, 22, 0, 0, 0, 0, time.Local),
				gid: 1,
				uid: 1,
			},
			wantTodayFirstCheckin: false,
			wantErr:               false,
		},
		{
			name: "测试非正常签到-3",
			args: args{
				ctx: ctx,
				t:   time.Date(2022, 3, 20, 0, 0, 0, 0, time.Local),
				gid: 1,
				uid: 1,
			},
			wantTodayFirstCheckin: false,
			wantErr:               false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTodayFirstCheckin, err := dao.Checkin(tt.args.ctx, tt.args.t, tt.args.gid, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Checkin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTodayFirstCheckin != tt.wantTodayFirstCheckin {
				t.Errorf("Checkin() gotTodayFirstCheckin = %v, want %v", gotTodayFirstCheckin, tt.wantTodayFirstCheckin)
			}
		})
	}
}

func TestDao_ListUserCheckinInfos(t *testing.T) {
	// TODO 完善单测
	dao := getDao()
	ctx := context.Background()

	checkins, err := dao.ListUserCheckinInfos(ctx, 1, 1)
	if err != nil {
		panic(err)
	}

	for _, item := range checkins {
		fmt.Printf("%#v\n", item)
	}
}
