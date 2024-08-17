// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"time"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Document struct {
// 	ActivityID int             `bson:"activityId" json:"activityId"`
// 	Nodes      json.RawMessage `bson:"nodes" json:"nodes"`
// 	Edges      json.RawMessage `bson:"edges" json:"edges"`
// }

// func main() {
// 	mongoURI := "mongodb://root:example@172.29.3.110:27017"
// 	dbName := "resources"
// 	collectionName := "activities"

// 	clientOptions := options.Client().ApplyURI(mongoURI)
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to MongoDB: %v", err)
// 	}
// 	defer func() {
// 		if err = client.Disconnect(context.TODO()); err != nil {
// 			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
// 		}
// 	}()

// 	collection := client.Database(dbName).Collection(collectionName)

// 	for i := 0; i < 300000; i++ {
// 		document := generateRandomDocument(i + 1)
// 		insertResult, err := collection.InsertOne(context.TODO(), document)
// 		if err != nil {
// 			log.Fatalf("Failed to insert document: %v", err)
// 		}
// 		if i % 10000 == 0{
// 			fmt.Printf("Inserted document with ID: %v\n", insertResult.InsertedID)
// 		}
// 	}
// }

// func generateRandomDocument(activityId int) Document {
// 	nodes := generateRandomJSON()
// 	edges := generateRandomJSON()

// 	return Document{
// 		ActivityID: activityId,
// 		Nodes:      nodes,
// 		Edges:      edges,
// 	}
// }

// func generateRandomJSON() json.RawMessage {
// 	randomData := map[string]interface{}{
// 		"randomValue": rand.Intn(100),
// 		"timestamp":   time.Now().Unix(),
// 		"details": map[string]interface{}{
// 			"key1": rand.Intn(1000),
// 			"key2": fmt.Sprintf("value%d", rand.Intn(100)),
// 		},
// 	}

// 	jsonData, err := json.Marshal(randomData)
// 	if err != nil {
// 		log.Fatalf("Failed to marshal JSON: %v", err)
// 	}

// 	return jsonData
// }

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

	// Number of documents to insert
	numDocuments := 5000000
	// Number of workers (Go routines)
	numWorkers := 1000
	// Create a channel to send documents to be inserted
	documentChan := make(chan Document, numWorkers)

	// Start worker Go routines
	for i := 0; i < numWorkers; i++ {
		go worker(collection, documentChan)
	}

	// Generate documents and send them to the workers
	for i := 0; i < numDocuments; i++ {
		document := generateRandomDocument(i + 1)
		documentChan <- document
		if i%100000 == 0 {
			fmt.Printf("Generated document #%d\n", i+1)
		}
	}

	// Close the channel to indicate that no more documents will be sent
	close(documentChan)
}

func worker(collection *mongo.Collection, documentChan <-chan Document) {
	for document := range documentChan {
		_, err := collection.InsertOne(context.TODO(), document)
		if err != nil {
			log.Printf("Failed to insert document: %v", err)
		} else {
			// fmt.Printf("Inserted document with ID: %v\n", insertResult.InsertedID)
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
