package mongo

import (
	"github.com/thylong/mongo_mock_go_example/model"
	"gopkg.in/mgo.v2"
)

func Connection() model.Session {
	mgoSession, err := mgo.Dial("mongodb://admin:admin123@127.0.0.1:27017")

	if err != nil {
		panic(err)
	}
	//documents, _ := session.DB("test").C("other_test").GetMyDocuments()
	return model.MongoSession{Session: mgoSession}
}

//func NewConnection() (*mongo.Database, )
