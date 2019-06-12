package api

import (
	"github.com/HichuYamichu/stream-app-server/app"
	"github.com/gorilla/mux"
)

// GetRouter : returns app router
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/videos/upload", app.UploadVideo).Methods("POST")
	router.HandleFunc("/api/video/{id}", app.ServeVideo).Methods("GET")
	router.HandleFunc("/api/videos", app.ListVideos).Methods("GET")
	router.HandleFunc("/api/images/miniatures/{id}", app.ServeMiniature).Methods("GET")
	return router
}
