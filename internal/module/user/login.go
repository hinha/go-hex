package user

import (
	"github.com/sirupsen/logrus"
	"testHEX/internal/constants/state"
)

func (s *service) Login(email, password string) (string, error) {
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

	user, err := s.userPersistence.Find(email, password)
	if err != nil {
		stringOperation = "persistence"
		return "0", err
	}

	//save to caching
	err = s.userCaching.Save(user)
	if err != nil {
		stringOperation = "cache"
	}
	return user.ID, err
}
