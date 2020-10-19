package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var(
	collection *mongo.Collection
	ctx = context.TODO()
)

func ConnectDB() *mongo.Client{
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil{
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil{
		log.Fatal(err)
	}
	log.Println("Database Successfully Configured!")
	return client
}