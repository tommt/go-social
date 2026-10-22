package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func Connect() {

	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	mongoURI := "mongodb://localhost:27017"
	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))

	if err == nil {
		fmt.Println("DATABASE CONNECTED")
		DB = client.Database("social")

	} else {
		fmt.Printf("error connecting to DB: %s\n", err.Error())
		// panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

}
