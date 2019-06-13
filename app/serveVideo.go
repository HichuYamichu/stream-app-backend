package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ServeVideo : serves a video to the client
func ServeVideo(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	location := "./store/videos/" + params["id"] + ".mp4"
	http.ServeFile(res, req, location)
}
