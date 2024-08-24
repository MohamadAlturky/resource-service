package main

import (
    "context"
    "fmt"
    "log"
    
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
)

func main() {

    mongoURI := "mongodb://root:example@172.29.3.110:27017"
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

    indexModel := mongo.IndexModel{
        Keys:    bson.D{{Key: "activityId", Value: 1}},
    }

    indexName, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
    if err != nil {
        log.Fatalf("Failed to create index: %v", err)
    }

    fmt.Printf("Index created: %s\n", indexName)
}
