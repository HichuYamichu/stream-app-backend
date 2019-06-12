package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Video : represents a video
type Video struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id"`
	Title string             `bson:"title" json:"title"`
	Desc  string             `bson:"desc" json:"desc"`
}
