package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	Collection    *mongo.Collection
)

func InitMongoDB() {

	uri := os.Getenv("MONGO_URI")


	clientOptions := options.Client().ApplyURI(uri)


	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}


	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}


	db := Client.Database(os.Getenv("MONGO_DB"))
	Collection = db.Collection(os.Getenv("MONGO_COLLECTION"))
}
