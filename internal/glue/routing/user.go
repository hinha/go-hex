package routing

import (
	"net/http"
	"testHEX/internal/module/user"
	"testHEX/platform/routers"
)

type userHandlers struct {
	handler user.Handler
}

// UserInit is to initialize the routers for domain user
func UserInit(handler user.Handler) user.Route {
	return &userHandlers{
		handler: handler,
	}
}

func (uh *userHandlers) Routers() []*routers.Router {
	return []*routers.Router{
		&routers.Router{
			Method:  http.MethodGet,
			URL:     "/test",
			Handler: uh.handler.Test,
		},
		&routers.Router{
			Method:  http.MethodPost,
			URL:     "/account/register",
			Handler: uh.handler.CreateNewAccount,
		},
		&routers.Router{
			Method: http.MethodPost,
			URL: "/account/login",
			Handler: uh.handler.SignIn,
		},
	}
}
