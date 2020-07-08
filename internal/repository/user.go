package repository

import (
	"testHEX/internal/constants/model"
	"testHEX/internal/module/user"
)

type userRepo struct {
	cache       user.Caching
	persistence user.Persistence
}

func (u userRepo) DataProfile(userID int64) (*model.User, error) {
	panic("implement me")
}

// UserInit to initiate the repository of user domain
func UserInit(cache user.Caching, persistence user.Persistence) user.Repository {
	return &userRepo{
		cache:       cache,
		persistence: persistence,
	}
}
