package app

import (
	"encoding/json"
	"net/http"

	"github.com/HichuYamichu/stream-app-server/db"
)

// ListVideos : lists all avaible videos
func ListVideos(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	videos := db.RetriveVideos()
	json.NewEncoder(res).Encode(videos)
}
