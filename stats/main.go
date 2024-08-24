package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := "mongodb://root:example@localhost:27017"
	dbName := "resources"
	collectionName := "activities"

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	collection := client.Database(dbName).Collection(collectionName)

	count, err := collection.CountDocuments(context.TODO(), map[string]interface{}{})
	if err != nil {
		log.Fatalf("Failed to count documents: %v", err)
	}

	fmt.Printf("The '%s' collection contains %d documents.\n", collectionName, count)
}
