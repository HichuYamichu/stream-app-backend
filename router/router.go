package router

import (
	"github.com/HichuYamichu/stream-app-server/controllers/dispatcher"
	"github.com/HichuYamichu/stream-app-server/controllers/videos"
	"github.com/gorilla/mux"
)

func Get() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/video/{name}", dispatcher.SendVideo).Methods("GET")
	router.HandleFunc("/api/videos/upload", videos.Upload).Methods("POST")
	router.HandleFunc("/api/videos", videos.ListAllVideos).Methods("GET")
	router.HandleFunc("/api/images/miniatures/{name}", videos.ServeVideoMiniatures).Methods("GET")
	return router
}
