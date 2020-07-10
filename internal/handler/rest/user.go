package rest

import (
	"encoding/json"
	"github.com/savsgio/atreugo/v10"
	"net/http"
	"testHEX/internal/constants/model"
	"testHEX/internal/module/security"
	"testHEX/internal/module/user"
)

type userService struct {
	usecase user.Usecase
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

	password := security.GeneratePasswordHash([]byte(reqObj.Password))

	err := us.usecase.Register(&model.User{
		Username: reqObj.Username,
		Email:    reqObj.Email,
		Password: password,
	})

	if err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		return ctx.JSONResponse(atreugo.JSON{
			"data": atreugo.JSON{
				"message": "Email already exists",
			},
			"status": "fail",
		})
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

	//us.hash.Write([]byte(reqObj.Password))
	//password := fmt.Sprintf("%x", us.hash.Sum(nil))
	//password := security.GeneratePasswordHash([]byte(reqObj.Password))
	_, token, err := us.usecase.Login(reqObj.Email, reqObj.Password)
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
			"token": token,
		},
	}, http.StatusOK)
}

// HandleUser is to initialize the rest handler for domain user
func HandleUser(usecase user.Usecase) user.Handler {
	return &userService{
		usecase: usecase,
	}
}

// Test is the test handler function
func (us *userService) Test(ctx *atreugo.RequestCtx) error {
	return ctx.TextResponse("Hello World")
}
