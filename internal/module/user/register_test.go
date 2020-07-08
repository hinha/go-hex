package user

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"testHEX/internal/constants/model"
	mock_user "testHEX/mocks/user"
	"testing"
)

func Test_service_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() (Persistence, Caching)
	}{
		{
			name: "error persistence create",
			args: args{
				user: &model.User{
					Username: "test_username",
					Email:    "test@email.com",
					Password: "hashedpassword",
				},
			},
			wantErr: true,
			initMock: func() (Persistence, Caching) {
				mockedPersis := mock_user.NewMockPersistence(ctrl)
				mockedPersis.EXPECT().Create(&model.User{
					Username: "test_username",
					Email:    "test@email.com",
					Password: "hashedpassword",
				}).Return(nil, errors.New("ERROR"))
				return mockedPersis, nil
			},
		},
		{
			name: "success",
			args: args{
				user: &model.User{
					Username: "test_username",
					Email:    "test@email.com",
					Password: "hashedpassword",
				},
			},
			initMock: func() (Persistence, Caching) {
				mockedPersis := mock_user.NewMockPersistence(ctrl)
				mockedPersis.EXPECT().Create(&model.User{
					Username: "test_username",
					Email:    "test@email.com",
					Password: "hashedpassword",
				}).Return(&model.User{
					ID:       "1",
					Username: "test_username",
					Email:    "test@email.com",
					Password: "hashedpassword",
				}, nil)
				mockedCache := mock_user.NewMockCaching(ctrl)
				mockedCache.EXPECT().Save(&model.User{
					ID:       "1",
					Username: "test_username",
					Email:    "test@email.com",
					Password: "hashedpassword",
				}).Return(errors.New("ERROR"))
				return mockedPersis, mockedCache
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, c := tt.initMock()
			s := &service{
				userPersistence: p,
				userCaching:     c,
			}
			if err := s.Register(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("service.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
