package presistence

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"strconv"
	"testHEX/internal/constants/model"
	"testHEX/internal/constants/state"
	"testHEX/internal/module/security"
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

func (u userPersistence) Find(email, password string) (*model.User, *model.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	users := new(model.User)
	token := new(model.Token)

	filter := bson.M{"email": email,"status":   state.UserActiveAccount}
	err := u.db.Collection(TABLE).FindOne(ctx, filter).Decode(users)
	if err != nil {
		return nil, nil, err
	}

	valid := security.PasswordCompare([]byte(password), []byte(users.Password))
	if valid != nil {
		return nil, nil, errors.New("LOGIN FAILED")
	}

	times := strconv.Itoa(int(time.Now().Unix()))
	yu := fmt.Sprintf("%s:%s:%s", times, users.ID, users.Email)
	enc, _ := security.EncryptString(yu, "ABCDEFG")
	token.UniqueToken = enc
	token.TimeAt = times

	return users, token, nil
}

func (u userPersistence) FindByEmail(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	users := new(model.User)

	filter := bson.M{"email": email}
	err := u.db.Collection(TABLE).FindOne(ctx, filter).Decode(users)

	if err == nil {
		return errors.New("USER EXISTS")
	}
	return nil
}

// UserInit is to init the user persistence that contains data accounts
func UserInit(db *mongo.Database) user.Persistence {
	return &userPersistence{
		db,
	}
}
