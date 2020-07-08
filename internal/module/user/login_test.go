package user

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
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
				email:    "clyf@email.com",
				password: "hashedpassword",
			},
			wantErr: true,
			initMock: func() (Persistence, Caching) {
				mockedPersis := mock_user.NewMockPersistence(ctrl)
				mockedPersis.EXPECT().Find("clyf@email.com", "hashedpassword").Return(nil, errors.New("ERROR"))
				return mockedPersis, nil
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
			got, err := s.Login(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
			fmt.Println(tt.want)
			if got != tt.want {
				t.Errorf("service.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}