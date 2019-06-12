package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertVideo : insetrs video data into db
func InsertVideo(title, desc string) string {
	i := primitive.NewObjectID()
	_, err := VideoColl.InsertOne(CTX, bson.M{
		"_id":   i,
		"title": title,
		"desc":  desc,
	})
	if err != nil {
		log.Fatal(err)
	}
	videoName := i.Hex()
	return videoName
}
