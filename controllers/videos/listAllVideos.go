package videos

import (
	"encoding/json"
	"fmt"
	"github.com/HichuYamichu/stream-app-server/storage"
	"github.com/HichuYamichu/stream-app-server/structs"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func ListAllVideos(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")

	var videos []structs.Video
	filter := bson.M{}
	cursor, err := storage.DB.Collection("videos").Find(ctx, filter)
	if err != nil {
		handleErr(err, res)
		return
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var video structs.Video
		cursor.Decode(&video)
		videos = append(videos, video)
		fmt.Println(video)
	}
	if err != nil {
		handleErr(err, res)
		return
	}
	json.NewEncoder(res).Encode(videos)
}
