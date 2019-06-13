package router

import (
	"net/http"

	"github.com/HichuYamichu/stream-app-server/app"
	"github.com/gorilla/mux"
)

// GetRouter : returns app router
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/videos/upload", app.UploadVideo).Methods("POST")
	api.HandleFunc("/video/{id}", app.ServeVideo).Methods("GET")
	api.HandleFunc("/videos", app.ListVideos).Methods("GET")
	r.PathPrefix("/miniatures/").Handler(http.StripPrefix("/miniatures/", http.FileServer(http.Dir("./store/miniatures"))))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/dist/")))
	r.PathPrefix("/").HandlerFunc(indexHandler("./web/dist/index.html"))
	return r
}

func indexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}
