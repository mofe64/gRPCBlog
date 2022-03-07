package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var DATABASE *mongo.Client = ConnectDB()

func ConnectDB() *mongo.Client {
	fmt.Println("Connecting to MongoDB...")
	client, clientErr := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if clientErr != nil {
		log.Fatalf("Error creating Mongodb client %v\n", clientErr)
	}
	connectionError := client.Connect(context.TODO())
	if connectionError != nil {
		log.Fatalf("Error connecting to mongodb client %v\n", clientErr)
	}
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error %v\n", err)
	}
	fmt.Println("Connected to MongoDB... ")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("goLangGrpcBlog").Collection(collectionName)
	return collection
}
