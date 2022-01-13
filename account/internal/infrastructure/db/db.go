package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27035")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Printf("%s", err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Connected to Database!")

	return client
}
