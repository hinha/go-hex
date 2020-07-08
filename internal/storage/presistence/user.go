package presistence

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"strconv"
	"testHEX/internal/constants/model"
	"testHEX/internal/constants/state"
	"testHEX/internal/module/user"
	"time"
)

type userPersistence struct {
	db *mongo.Database
}

const (
	TIMEOUT = 30 * time.Second
	TABLE   = "users"
)

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func (u userPersistence) Create(user *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	rand.Seed(time.Now().UnixNano())
	ids := strconv.Itoa(randomInt(1000000000, 9999999999))
	user.CreatedAt = time.Now().Unix()
	user.ID = ids
	//fmt.Println(user)
	
	if u.db.Client() == nil {
		return nil, nil
	}
	_, err := u.db.Collection(TABLE).InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userPersistence) FindByID(userID int64) (*model.User, error) {
	panic("implement me")
}

func (u userPersistence) Find(email, password string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	users := new(model.User)
	fmt.Println(password)
	filter := bson.M{"email": email, "password": password, "status":   state.UserActiveAccount}
	err := u.db.Collection(TABLE).FindOne(ctx, filter).Decode(users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// UserInit is to init the user persistence that contains data accounts
func UserInit(db *mongo.Database) user.Persistence {
	return &userPersistence{
		db,
	}
}
