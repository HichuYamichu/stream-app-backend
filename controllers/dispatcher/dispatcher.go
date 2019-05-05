package dispatcher

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SendVideo(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	location := "g:/Videos/testing/" + params["name"] + ".mkv"
	http.ServeFile(res, req, location)
}
