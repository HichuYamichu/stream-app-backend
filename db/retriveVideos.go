package db

import (
	"log"

	"github.com/HichuYamichu/stream-app-server/models"

	"go.mongodb.org/mongo-driver/bson"
)

// RetriveVideos : returns all vidoes data
func RetriveVideos() []models.Video {
	var videos []models.Video
	filter := bson.M{}
	cursor, err := VideoColl.Find(CTX, filter)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(CTX)
	for cursor.Next(CTX) {
		var video models.Video
		cursor.Decode(&video)
		videos = append(videos, video)
	}
	if err != nil {
		log.Fatal(err)
	}
	return videos
}
