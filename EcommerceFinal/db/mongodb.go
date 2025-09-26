package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(""))
	if err != nil {
		log.Fatal(err)
	}
	return client

}
