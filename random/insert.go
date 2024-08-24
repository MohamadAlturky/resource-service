package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Document struct {
	ActivityID int             `bson:"activityId" json:"activityId"`
	Nodes      json.RawMessage `bson:"nodes" json:"nodes"`
	Edges      json.RawMessage `bson:"edges" json:"edges"`
}

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

	numDocuments := 10000000

	
	numWorkers := 100000
	
	
	documentChan := make(chan Document, numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(collection, documentChan)
	}

	for i := 0; i < numDocuments; i++ {
		document := generateRandomDocument(i + 1)
		documentChan <- document
		if i%100000 == 0 {
			fmt.Printf("Generated document #%d\n", i+1)
		}
	}

	close(documentChan)
}

func worker(collection *mongo.Collection, documentChan <-chan Document) {
	for document := range documentChan {
		_, err := collection.InsertOne(context.TODO(), document)
		if err != nil {
			log.Printf("Failed to insert document: %v", err)
		} else {
		}
	}
}

func generateRandomDocument(activityId int) Document {
	nodes := generateRandomJSON()
	edges := generateRandomJSON()

	return Document{
		ActivityID: activityId,
		Nodes:      nodes,
		Edges:      edges,
	}
}

func generateRandomJSON() json.RawMessage {
	randomData := map[string]interface{}{
		"randomValue": rand.Intn(100),
		"timestamp":   time.Now().Unix(),
		"details": map[string]interface{}{
			"key1": rand.Intn(1000),
			"key2": fmt.Sprintf("value%d", rand.Intn(100)),
		},
	}

	jsonData, err := json.Marshal(randomData)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	return jsonData
}
