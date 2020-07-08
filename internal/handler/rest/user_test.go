package rest

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/savsgio/atreugo/v10"
	"github.com/valyala/fasthttp"
	"hash"
	"net/http"
	"reflect"
	"strings"
	"testHEX/internal/constants/model"
	"testHEX/internal/module/user"
	mockuser "testHEX/mocks/user"
	"testing"
)

func TestHandleUser(t *testing.T) {
	type args struct {
		usecase user.Usecase
		hash    hash.Hash
	}

	tests := []struct {
		name string
		args args
		want user.Handler
	}{
		{
			name: "success",
			args: args{
				usecase: nil,
				hash:    nil,
			},
			want: &userService{
				usecase: nil,
				hash:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleUser(tt.args.usecase, tt.args.hash); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Test(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name       string
		args       args
		want       string
		wantStatus int
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{},
				},
			},
			want:       "Hello World",
			wantStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &responseBodyWriter{body: bytes.NewBufferString("Hello World"), ResponseWriter: tt.args.ctx.Writer, statusCode: 200}
			got := w.body.String()
			if got != tt.want {
				t.Errorf("Response.Body() = %v, want %v", got, tt.want)
			}
			gotStatus := w.statusCode
			if gotStatus != tt.wantStatus {
				t.Errorf("Response.StatusCode() = %v, want %v", got, tt.wantStatus)
			}
		})
	}
}

func Test_userService_CreateNewAccount(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		want       interface{}
		wantStatus int
		keyCookie  string
		wantCookie string
		wantErr    bool
		initMock   func() (*atreugo.RequestCtx, user.Usecase)
	}{
		{
			name: "bad username",
			want: atreugo.JSON{
				"data": atreugo.JSON{
					"message": "missing required fields",
				},
				"status": http.StatusText(http.StatusBadRequest),
			},
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
			initMock: func() (*atreugo.RequestCtx, user.Usecase) {
				ctx := &atreugo.RequestCtx{
					RequestCtx: &fasthttp.RequestCtx{},
				}
				ctx.Request.PostArgs().Add("email", "test@email.com")
				ctx.Request.PostArgs().Add("password", "nothashedpassword")
				return ctx, nil
			},
		},
		{
			name: "bad email",
			want: atreugo.JSON{
				"data": atreugo.JSON{
					"message": "missing required fields",
				},
				"status": http.StatusText(http.StatusBadRequest),
			},
			wantStatus: 400,
			wantErr:    true,
			initMock: func() (*atreugo.RequestCtx, user.Usecase) {
				ctx := &atreugo.RequestCtx{
					RequestCtx: &fasthttp.RequestCtx{},
				}
				ctx.Request.PostArgs().Add("username", "test")
				ctx.Request.PostArgs().Add("password", "nothashedpassword")
				return ctx, nil
			},
		},
		{
			name: "bad password",
			want: atreugo.JSON{
				"data": atreugo.JSON{
					"message": "missing required fields",
				},
				"status": http.StatusText(http.StatusBadRequest),
			},
			wantStatus: 400,
			wantErr:    true,
			initMock: func() (*atreugo.RequestCtx, user.Usecase) {
				ctx := &atreugo.RequestCtx{
					RequestCtx: &fasthttp.RequestCtx{},
				}
				ctx.Request.PostArgs().Add("username", "usertest")
				ctx.Request.PostArgs().Add("email", "clyf@email.com")
				return ctx, nil
			},
		},
		{
			name: "error usecase register",
			want: atreugo.JSON{
				"email": "clyf@email.com",
			},
			wantStatus: 201,
			wantErr:    true,
			initMock: func() (*atreugo.RequestCtx, user.Usecase) {
				ctx := &atreugo.RequestCtx{
					RequestCtx: &fasthttp.RequestCtx{},
				}

				ctx.Request.SetBody([]byte(`{
					"username": "usertest",
					"email": "clyf@email.com",
					"password": "nothashedpassword"
				}`))

				mockedUsecase := mockuser.NewMockUsecase(ctrl)
				mockedUsecase.EXPECT().Register(&model.User{
					Username: "usertest",
					Email:    "clyf@email.com",
					Password: "ae7bc8bd67f4f5c6a911373677ec56f95246288f1130c62048703be38397bda7",
				}).Return(errors.New("ERROR"))

				return ctx, mockedUsecase
			},
		},
		{
			name: "success",
			want: atreugo.JSON{
				"email": "clyf@email.com",
			},
			wantStatus: 201,
			wantErr:    true,
			initMock: func() (*atreugo.RequestCtx, user.Usecase) {
				ctx := &atreugo.RequestCtx{
					RequestCtx: &fasthttp.RequestCtx{},
				}

				ctx.Request.SetBody([]byte(`{
					"username": "usertest",
					"email": "clyf@email.com",
					"password": "nothashedpassword"
				}`))

				mockedUsecase := mockuser.NewMockUsecase(ctrl)
				mockedUsecase.EXPECT().Register(&model.User{
					Username: "usertest",
					Email:    "clyf@email.com",
					Password: "ae7bc8bd67f4f5c6a911373677ec56f95246288f1130c62048703be38397bda7",
				}).Return(errors.New("ERROR"))

				return ctx, mockedUsecase
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, usecase := tt.initMock()
			us := &userService{
				hash:    sha256.New(),
				usecase: usecase,
			}

			if err := us.CreateNewAccount(ctx); (err == nil) != tt.wantErr {
				t.Errorf("userService.CreateNewAccount() error = %v, wantErr %v", err, tt.wantErr)
			}

			jsonString, err := json.Marshal(tt.want)
			if err != nil {
				fmt.Println(err)
			}
			tt.want = string(jsonString)
			got := string(ctx.Response.Body())

			if got != tt.want {
				t.Errorf("Response.Body() = %v, want %v", got, tt.want)
			}
			gotStatus := ctx.Response.StatusCode()
			if gotStatus != tt.wantStatus {
				t.Errorf("Response.StatusCode() = %v, want %v", gotStatus, tt.wantStatus)
			}
			if tt.keyCookie == "" {
				return
			}
			gotCookie := string(ctx.Response.Header.PeekCookie(tt.keyCookie))
			if !strings.Contains(got, fmt.Sprint("=", tt.wantCookie)) && tt.wantCookie != "" {
				t.Errorf("Response.Body() = %v, want %v", gotCookie, tt.wantCookie)
			}
		})
	}

}

//func typeof(v interface{}) string {
//	return fmt.Sprintf("%T", v)
//}
