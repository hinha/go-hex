package presistence

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testHEX/internal/constants/model"
	"testHEX/internal/module/user"
	mymong "testHEX/platform/mongo"
	"testing"
)

const (
	mongoURL = "mongodb://admin:admin123@127.0.0.1:27017"
	mongoDB  = "tests"
)

func TestUserInit(t *testing.T) {
	type args struct {
		db *mongo.Database
	}

	tests := []struct {
		name string
		args args
		want user.Persistence
	}{
		{
			name: "success",
			args: args{
				db: nil,
			},
			want: &userPersistence{
				db: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserInit(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				fmt.Printf("UserInit() = %v, want %v", got, tt.want)
				t.Errorf("UserInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userPersistence_Create(t *testing.T) {
	//documents, _ := session.DB("test").C("other_test").GetMyDocuments()

	dbs := mymong.Connection(mongoURL, mongoDB)

	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name     string
		args     args
		want     *model.User
		wantErr  bool
		initMock func() *mongo.Database
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				user: &model.User{
					Username: "clyf",
					Email:    "clyf@email.com",
					Password: "hashedpassword",
				},
			},
			want: &model.User{
				ID:       "1",
				Username: "clyf",
				Email:    "clyf@email.com",
				Password: "hashedpassword",
			},
			initMock: func() *mongo.Database {
				ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
				defer cancel()
				_, _ = dbs.Collection("users").InsertOne(ctx, &model.User{
					Username: "clyf",
					Email:    "clyf@email.com",
					Password: "hashedpassword",
				})
				return &mongo.Database{}
			},
		},
	}
	//
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//fmt.Println(tt.initMock())
			up := &userPersistence{
				db: tt.initMock(),
			}
			//fmt.Println(tt.args.ctx)
			//fmt.Println(tt.args.user)
			//fmt.Println(tt.wantErr)
			got, _ := up.Create(tt.args.user)
			//if (err != nil) == tt.wantErr {
			//	t.Errorf("userPersistence.Create() error = %v, wantErr %v", err, tt.wantErr)
			//	return
			//}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("userPersistence.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
