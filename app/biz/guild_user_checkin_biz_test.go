package biz_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/hkail/taskbot/app/biz"

	"github.com/hkail/taskbot/app/biz/dao"
)

func TestGetUserCheckinDateStatusMapper(t *testing.T) {
	type args struct {
		userCheckins []dao.GuildUserCheckin
	}
	tests := []struct {
		name string
		args args
		want map[int]bool
	}{
		{
			name: "测试获取用户签到信息是否正常",
			args: args{
				userCheckins: []dao.GuildUserCheckin{
					{
						Yearmonth: 202203,
						Days:      23, // ...00010111
					},
				},
			},
			want: map[int]bool{
				20220301: true,
				20220302: true,
				20220303: true,
				20220304: false,
				20220305: true,
				20220306: false,
				20220307: false,
				20220308: false,
				20220309: false,
				20220310: false,
				20220311: false,
				20220312: false,
				20220313: false,
				20220314: false,
				20220315: false,
				20220316: false,
				20220317: false,
				20220318: false,
				20220319: false,
				20220320: false,
				20220321: false,
				20220322: false,
				20220323: false,
				20220324: false,
				20220325: false,
				20220326: false,
				20220327: false,
				20220328: false,
				20220329: false,
				20220330: false,
				20220331: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biz.GetUserCheckinDateStatusMapper(tt.args.userCheckins); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserCheckinDateStatusMapper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBiz_GetUserContCheckinDays(t *testing.T) {
	biz := getBiz()
	ctx := context.Background()

	type args struct {
		ctx context.Context
		gid uint64
		uid uint64
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "测试获取连续签到天数-1",
			args: args{
				ctx: ctx,
				gid: 1,
				uid: 1,
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := biz.GetUserContCheckinDays(tt.args.ctx, tt.args.gid, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserContCheckinDays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserContCheckinDays() got = %v, want %v", got, tt.want)
			}
		})
	}
}
