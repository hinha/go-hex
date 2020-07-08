package cache

import (
	"github.com/go-redis/redis"
	"reflect"
	"testHEX/internal/constants/model"
	"testHEX/internal/module/user"
	mockredis "testHEX/mocks/redis"
	"testing"
	"time"
)

func TestUserInit(t *testing.T) {
	type args struct {
		conn *redis.Client
	}
	tests := []struct {
		name string
		args args
		want user.Caching
	}{
		{
			name: "success",
			args: args{
				conn: nil,
			},
			want: &userCache{
				connection: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserInit(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userCache_Save(t *testing.T) {
	client, miniredis := mockredis.Connection()

	type args struct {
		user *model.User
	}

	tests := []struct {
		name      string
		args      args
		wantErr   bool
		expecting func()
	}{
		{
			name: "redis exists",
			args: args{
				user: &model.User{
					ID:        "1",
					Email:     "test@email.com",
					Username:  "test",
					CreatedAt: time.Date(2020, time.January, 12, 12, 12, 12, 12, time.UTC).Unix(),
					LastLogin: time.Date(2020, time.January, 12, 12, 12, 12, 12, time.UTC).Unix(),
					Status:    1,
				},
			},
			expecting: func() {
				miniredis.CheckGet(t, "user:1", `{"id":"1","username":"test","email":"test@email.com","created_at":1578831132,"last_login":1578831132,"status":1}`)

			},
		},
		{
			name: "redis expire",
			args: args{
				user: &model.User{
					ID:        "1",
					Email:     "test@email.com",
					Username:  "test",
					CreatedAt: time.Date(2020, time.January, 12, 12, 12, 12, 12, time.UTC).Unix(),
					LastLogin: time.Date(2020, time.January, 12, 12, 12, 12, 12, time.UTC).Unix(),
					Status:    1,
				},
			},
			expecting: func() {
				miniredis.FastForward(time.Second * 60 * 60 * 25)
				if miniredis.Exists("user:1") {
					t.Error("This should not be existed anymore")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &userCache{
				connection: client,
			}
			if err := uc.Save(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userCache.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.expecting()
		})
	}
	miniredis.Close()
}

func Test_userCache_Get(t *testing.T) {
	client, miniredis := mockredis.Connection()

	type args struct {
		userID string
	}
	tests := [] struct {
		name     string
		args     args
		want     *model.User
		wantErr  bool
		initMock func()
	} {
		{
			name: "not exist",
			args: args{
				userID: "1",
			},
			wantErr: true,
			initMock: func() {},
		},
		{
			name: "exists",
			args: args{
				userID: "1",
			},
			want: &model.User{
				ID:        "1",
				Email:     "test@email.com",
				Username:  "test",
				CreatedAt: time.Date(2020, time.January, 12, 12, 12, 12, 12, time.UTC).Unix(),
				LastLogin: time.Date(2020, time.January, 12, 12, 12, 12, 12, time.UTC).Unix(),
				Status:    1,
			},
			initMock: func() {
				_ = miniredis.Set("user:1", `{"id":"1","username":"test","email":"test@email.com","created_at":1578831132,"last_login":1578831132,"status":1}`)
				miniredis.SetTTL("user:1", time.Second*60*60*24)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initMock()
			uc := &userCache{
				connection: client,
			}
			got, err := uc.Get(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userCache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
	miniredis.Close()
}

func Test_userCache_Delete(t *testing.T) {
	client, miniredis := mockredis.Connection()

	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				userID: "1",
			},
			wantErr: miniredis.Exists("user:1"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &userCache{
				connection: client,
			}
			if err := uc.Delete(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("userCache.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	miniredis.Close()
}