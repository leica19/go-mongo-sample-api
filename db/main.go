package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://db")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func CreateCollection(client *mongo.Client) {
	client.Database("test").Collection("users")
	log.Print("** new collection created **")
}

func GetConnectionOfCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("test").Collection("users")
	return collection
}
