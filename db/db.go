package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// VideoColl : Videos collection
var VideoColl *mongo.Collection

// CTX : db context
var CTX, _ = context.WithTimeout(context.Background(), 30*time.Second)

// Connect : conntect to MongoDB instance
func Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(CTX)
	if err != nil {
		log.Fatal(err)
	}
	VideoColl = client.Database("stream-app").Collection("videos")
}
