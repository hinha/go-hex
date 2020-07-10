package user

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testHEX/internal/constants/state"
)

func (s *service) Login(email, password string) (string, string, error) {
	var err error
	var stringOperation string

	defer func() {
		if err == nil {
			return
		}

		// log formatting
		log := logrus.WithFields(logrus.Fields{
			"domain":  "user",
			"action":  "login user",
			"usecase": "Login",
		})
		log.WithField(state.LogType, stringOperation).Errorln(err)
	}()
	user, token, err := s.userPersistence.Find(email, password)
	fmt.Println("is error", err)
	fmt.Println("user error", user)
	if err != nil {
		stringOperation = "persistence"
		return "0", "-", err
	}

	//times := strconv.Itoa(int(time.Now().Unix()))
	//yu := fmt.Sprintf("%s:%s:%s", times, user.ID, user.Email)
	//enc, _ := security.EncryptString(yu, "ABCDEFG")

	err = s.userCaching.SaveToken(token, user)
	if err != nil {
		stringOperation = "cache"
	}
	fmt.Println("error caching 1: ", err)
	//dec, err := security.DecryptString("Nwh1vl9Z27sxM88Hz5NV47TnQJkvyWBwC3Ru8V1ybS4BvQDVhyz1YT1UMgrFse4y7jEXMuOo", "ABCDEFG")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(dec)

	//save to caching
	err = s.userCaching.Save(user)
	if err != nil {
		stringOperation = "cache"
	}
	fmt.Println("error caching 2: ", err)
	return user.ID, token.UniqueToken, err
}
