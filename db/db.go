package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// VideoColl : Videos collection
var VideoColl *mongo.Collection

// CTX : db context
var CTX context.Context

// Connect : conntect to MongoDB instance
func Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	CTX, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(CTX)
	if err != nil {
		panic(err)
	}
	VideoColl = client.Database("stream-app").Collection("videos")
}
