package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Video struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id"`
	Title string             `bson:"title" json:"title"`
	Desc  string             `bson:"desc" json:"desc"`
}
