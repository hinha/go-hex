package user

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"testHEX/internal/constants/model"
	mock_user "testHEX/mocks/user"
	"testing"
)

func Test_service_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name     string
		args     args
		want     string
		wantErr  bool
		initMock func() (Persistence, Caching)
	} {
		{
			name: "error persistence find",
			args: args{
				email:    "test@email.com",
				password: "hashedpassword",
			},
			wantErr: true,
			initMock: func() (Persistence, Caching) {
				mockedPersis := mock_user.NewMockPersistence(ctrl)
				mockedPersis.EXPECT().Find("test@email.com", "hashedpassword").Return(nil, nil, errors.New("ERROR"))
				return mockedPersis, nil
			},
			want: "0",
		},
		{
			name: "error caching save",
			args: args{
				email:    "test@email.com",
				password: "hashedpassword",
			},
			want:    "1",
			wantErr: true,
			initMock: func() (Persistence, Caching) {
				mockedPersis := mock_user.NewMockPersistence(ctrl)
				mockedPersis.EXPECT().Find("test@email.com", "hashedpassword").Return(&model.User{
					ID:       "1",
					Username: "test",
					Email:    "test@email.com",
					Password: "hashedpassword",
				}, &model.Token{
					UniqueToken: "TGrGeaK5iXSFa_GSyeBtIVdpHBHiofg3lRTmUSVgK-alepTHC-25aFInYg==",
					TimeAt: "1594374962",
				}, nil)
				mockedCache := mock_user.NewMockCaching(ctrl)
				mockedCache.EXPECT().SaveToken(&model.Token{
					UniqueToken: "TGrGeaK5iXSFa_GSyeBtIVdpHBHiofg3lRTmUSVgK-alepTHC-25aFInYg==",
					TimeAt: "1594374962",
				},&model.User{
					ID:       "1",
					Username: "test",
					Email:    "test@email.com",
					Password: "hashedpassword",
				}).Return(errors.New("ERROR"))

				mockedCache.EXPECT().Save(&model.User{
					ID:       "1",
					Username: "test",
					Email:    "test@email.com",
					Password: "hashedpassword",
				}).Return(errors.New("ERROR"))
				fmt.Println(mockedCache, mockedPersis)
				return mockedPersis, mockedCache
			},
		},
		{
			name: "success",
			args: args{
				email:    "test@email.com",
				password: "hashedpassword",
			},
			want: "1",
			initMock: func() (Persistence, Caching) {
				mockedPersis := mock_user.NewMockPersistence(ctrl)
				mockedPersis.EXPECT().Find("test@email.com", "hashedpassword").Return(&model.User{
					ID:       "1",
					Username: "test",
					Email:    "test@email.com",
					Password: "hashedpassword",
				}, &model.Token{
					UniqueToken: "TGrGeaK5iXSFa_GSyeBtIVdpHBHiofg3lRTmUSVgK-alepTHC-25aFInYg==",
					TimeAt: "1594374962",
				}, nil)
				mockedCache := mock_user.NewMockCaching(ctrl)
				mockedCache.EXPECT().SaveToken(&model.Token{
					UniqueToken: "TGrGeaK5iXSFa_GSyeBtIVdpHBHiofg3lRTmUSVgK-alepTHC-25aFInYg==",
					TimeAt: "1594374962",
				},&model.User{
					ID:       "1",
					Username: "test",
					Email:    "test@email.com",
					Password: "hashedpassword",
				}).Return(nil)

				mockedCache.EXPECT().Save(&model.User{
					ID:       "1",
					Username: "test",
					Email:    "test@email.com",
					Password: "hashedpassword",
				}).Return(nil)
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
			got, _, err := s.Login(tt.args.email, tt.args.password)
			fmt.Println("got: ", got)
			fmt.Println("wantErr",  tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}