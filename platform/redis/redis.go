package redis

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type connectionString struct {
	connection string
	password   string
	domain     string
}

type Connections interface {
	Open() *redis.Client
}

func Initialize(connection, password, domain string) Connections {
	return &connectionString{
		connection: connection,
		password:   password,
		domain:     domain,
	}
}

func (cs *connectionString) Open() *redis.Client {
	logrus.WithFields(logrus.Fields{
		"platform": "redis",
		"domain":   cs.domain,
	}).Info("Connection to redis")

	client := redis.NewClient(&redis.Options{
		Addr:     cs.connection,
		Password: cs.password,
		DB:       1,
	})
	err := client.Ping().Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"connection": cs.connection,
			"password":   cs.password,
		}).Fatal(err)
		panic(err)
	}

	return client
}
