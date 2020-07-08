package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"testHEX/internal/constants/model"
	"testHEX/internal/module/user"
	"time"
)

type userCache struct {
	connection *redis.Client
}

const (
	keyRedisUser   = "user:%s"
	expireRedisKey = time.Second * 60 * 60 * 24
)

// UserInit is the function to init the user caching
func UserInit(conn *redis.Client) user.Caching {
	return &userCache{
		connection: conn,
	}
}

func (uc *userCache) Save(user *model.User) error {

	data, _ := json.Marshal(user)
	key := fmt.Sprintf(keyRedisUser, user.ID)
	err := uc.connection.Set(key, data, time.Duration(expireRedisKey)).Err()
	user.Password = "" // no sensitive data allowed to be saved in cache
	return err
}

func (uc *userCache) Get(userID string) (*model.User, error) {
	key := fmt.Sprintf(keyRedisUser, userID)

	res, err := uc.connection.Get(key).Result()
	if err != nil {
		return nil, err
	}

	userN := new(model.User)
	err = json.Unmarshal([]byte(res), userN)

	return userN, nil
}

func (uc *userCache) Delete(userID string) error {
	key := fmt.Sprintf(keyRedisUser, userID)
	err := uc.connection.Del(key).Err()

	return err
}
