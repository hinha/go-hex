package user

import (
	"github.com/sirupsen/logrus"
	"testHEX/internal/constants/model"
	"testHEX/internal/constants/state"
)

func (s *service) Register(user *model.User) error {
	// log formatting
	log := logrus.WithFields(logrus.Fields{
		"domain":  "user",
		"action":  "create new user",
		"usecase": "Register",
	})

	user, err := s.userPersistence.Create(user)
	if err != nil {
		log.WithField(state.LogType, "persistence").Errorln(err)
		return err
	}

	//this saving to the cache is a business logic
	//user can automatically login after registration is A business logic because it is a flow of user experience
	//so, don't put your business logic inside repository
	//repository is just a data storing logic
	err = s.userCaching.Save(user)
	if err != nil {
		log.WithField(state.LogType, "caching").Errorln(err)
	}
	//there's no need for user to know if the caching is failed
	return nil
}
