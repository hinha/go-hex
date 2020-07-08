package rest

import (
	"encoding/json"
	"fmt"
	"github.com/savsgio/atreugo/v10"
	"hash"
	"net/http"
	"testHEX/internal/constants/model"
	"testHEX/internal/module/user"
)

type userService struct {
	usecase user.Usecase
	hash    hash.Hash
}

func (us *userService) CreateNewAccount(ctx *atreugo.RequestCtx) error {
	reqObj := new(model.RequestRegister)
	body := ctx.PostBody()

	if json.Unmarshal(body, &reqObj) != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		return ctx.JSONResponse(atreugo.JSON{
			"data": atreugo.JSON{
				"message": "missing required fields",
			},
			"status": http.StatusText(http.StatusBadRequest),
		}, http.StatusBadRequest)
	}

	if reqObj.Email == "" || reqObj.Username == "" || reqObj.Password == "" {
		ctx.SetStatusCode(http.StatusBadRequest)
		return ctx.JSONResponse(atreugo.JSON{
			"data": atreugo.JSON{
				"message": "missing required fields",
			},
			"status": http.StatusText(http.StatusBadRequest),
		}, http.StatusBadRequest)
	}

	us.hash.Write([]byte(reqObj.Password))
	reqObj.Password = fmt.Sprintf("%x", us.hash.Sum(nil))

	err := us.usecase.Register(&model.User{
		Username: reqObj.Username,
		Email:    reqObj.Email,
		Password: reqObj.Password,
	})
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
	}

	return ctx.JSONResponse(atreugo.JSON{
		"email": reqObj.Email,
	}, http.StatusCreated)
}

func (us *userService) SignIn(ctx *atreugo.RequestCtx) error {
	reqObj := new(model.RequestLogin)
	body := ctx.PostBody()

	if json.Unmarshal(body, &reqObj) != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		return ctx.JSONResponse(atreugo.JSON{
			"data": atreugo.JSON{
				"message": "missing required fields",
			},
			"status": http.StatusText(http.StatusBadRequest),
		}, http.StatusBadRequest)
	}

	if reqObj.Email == "" || reqObj.Password == "" {
		ctx.SetStatusCode(http.StatusBadRequest)
		return ctx.JSONResponse(atreugo.JSON{
			"data": atreugo.JSON{
				"message": "missing required fields",
			},
			"status": http.StatusText(http.StatusBadRequest),
		}, http.StatusBadRequest)
	}

	us.hash.Write([]byte(reqObj.Password))
	password := fmt.Sprintf("%x", us.hash.Sum(nil))
	_, err := us.usecase.Login(reqObj.Email, password)
	if err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		return ctx.JSONResponse(atreugo.JSON{
			"status": "success",
			"data": atreugo.JSON{
				"message": "Email or password not valid",
			},
		}, http.StatusOK)
	}

	return ctx.JSONResponse(atreugo.JSON{
		"status": "success",
		"data": atreugo.JSON{
			"message": "Successfully user",
		},
	}, http.StatusOK)
}

// HandleUser is to initialize the rest handler for domain user
func HandleUser(usecase user.Usecase, hash hash.Hash) user.Handler {
	return &userService{
		usecase: usecase,
		hash:    hash,
	}
}

// Test is the test handler function
func (us *userService) Test(ctx *atreugo.RequestCtx) error {
	return ctx.TextResponse("Hello World")
}
