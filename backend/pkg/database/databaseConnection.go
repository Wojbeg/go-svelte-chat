package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBinstance()

func DBinstance() *mongo.Client {

	MongoDb := os.Getenv("DB_ADDRESS")

	if MongoDb == "" {
		log.Fatal("Empty database address")
	}

	fmt.Println("MongoDb address: ", MongoDb)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Failed to connect to mongodb")
		return nil
	}

	fmt.Println("Successfully connected to MongoDB")
	fmt.Println("Mongo address: ", MongoDb)
	return client
}

func MessagesData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Chat").Collection(collectionName)
	return collection
}
