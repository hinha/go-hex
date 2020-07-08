package user

import (
	"github.com/savsgio/atreugo/v10"
	"testHEX/internal/constants/model"
	"testHEX/platform/routers"
)

type Persistence interface {
	Create(user *model.User) (*model.User, error)
	FindByID(userID int64) (*model.User, error)
	Find(email, password string) (*model.User, error)
	FindByEmail(email string) error
	//ChangePassword(newPassword string, user *model.User) error
	//Delete(user *model.User) error
}

// Caching initiator contains functions to get data from redis
type Caching interface {
	Save(user *model.User) error
	Get(userID string) (*model.User, error)
	Delete(userID string) error
}

type Repository interface {
	DataProfile(userID int64) (*model.User, error)
}

type service struct {
	userPersistence Persistence
	userCaching     Caching
	userRepository  Repository
}

type Usecase interface {
	Login(email, password string) (string, error)
	//Profile(userID int64) (*model.User, error)
	Register(user *model.User) error
}

func InitializeDomain(persistence Persistence, caching Caching, repository Repository) Usecase {
	return &service{
		userPersistence: persistence,
		userCaching:     caching,
		userRepository:  repository,
	}
}

type Handler interface {
	Test(ctx *atreugo.RequestCtx) error
	CreateNewAccount(ctx *atreugo.RequestCtx) error
	SignIn(ctx *atreugo.RequestCtx) error
	//ShowProfile(ctx *atreugo.RequestCtx) error
}

// Route contains the functions that will be used for the routing domain user
type Route interface {
	Routers() []*routers.Router
}
