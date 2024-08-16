package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	col    *mongo.Collection
)

func main() {
	// Initialize MongoDB connection
	initMongoDB()

	// Create a new Gin router
	router := gin.Default()

	// Define routes
	router.POST("/set", setHandler)
	router.GET("/get/:activityId", getHandler)

	// Start the server
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initMongoDB() {
	// MongoDB connection URI
	uri := "mongodb://root:example@localhost:27017"

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	// Set database and collection
	db := client.Database("resources")
	col = db.Collection("activities")
}

// Handler for the /set endpoint
func setHandler(c *gin.Context) {
	var input map[string]interface{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	activityID, ok := input["activityId"].(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activityId"})
		return
	}

	// Convert activityID to int
	activityIDInt := int(activityID)

	// Insert document into MongoDB
	_, err := col.InsertOne(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"activityId": activityIDInt})
}

// Handler for the /get/:activityId endpoint
func getHandler(c *gin.Context) {
	activityIDStr := c.Param("activityId")
	activityID, err := strconv.Atoi(activityIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activityId"})
		return
	}

	// Find document in MongoDB
	filter := bson.M{"activityId": activityID}
	var result bson.M
	err = col.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find document"})
		}
		return
	}

	c.JSON(http.StatusOK, result)
}
