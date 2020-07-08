package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"testHEX/internal/glue/routing"
	"testHEX/internal/handler/rest"
	"testHEX/internal/module/user"
	"testHEX/internal/repository"
	"testHEX/internal/storage/cache"
	"testHEX/internal/storage/presistence"
	"testHEX/platform/mongo"
	"testHEX/platform/redis"
	"testHEX/platform/routers"
)

const (
	// change all connection strings based on your own credentials
	redisURL      = "127.0.0.1:6379"
	redisPassword = ""
	// mysqlURL   = "postgresql://postgres@127.0.0.1/postgres?sslmode=disable"
	mysqlURL = "tcp:root:alongside@localhost/test_user?sslmode=disable"

	domain = "user"

	mongoURL = "mongodb://admin:admin123@127.0.0.1:27017"
	mongoDB  = "tests"
)

var testInit bool

func main() {
	//MYmysql := mysql.Initialize(mysqlURL, mysqlURL, domain)
	//mysqlConn := MYmysql.Open()

	rds := redis.Initialize(redisURL, redisPassword, domain)
	redisConn := rds.Open()

	mongoConnection := mongo.Connection(mongoURL, mongoDB)
	database := presistence.UserInit(mongoConnection)
	caching := cache.UserInit(redisConn)
	repo := repository.UserInit(caching, database)

	usecase := user.InitializeDomain(database, caching, repo)

	handler := rest.HandleUser(usecase)
	router := routing.UserInit(handler).Routers()
	servant := routers.Initialize(":9000", router, domain)
	if testInit {
		logrus.Info("Initialize test mode Finished!")
		os.Exit(0)
	}

	servant.Serve()
}
