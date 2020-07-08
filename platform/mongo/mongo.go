package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connection(uri string, dbName string) *mongo.Database {
	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(uri).SetMaxPoolSize(10))

	if err != nil {
		panic(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		panic(err)
	}

	db := client.Database(dbName)
	fmt.Println("Connected to Mongo DB")

	return db
}
