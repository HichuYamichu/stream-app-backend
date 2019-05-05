package videos

import (
	"github.com/gorilla/mux"
	"net/http"
)

func ServeVideoMiniatures(res http.ResponseWriter, req *http.Request) {
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	params := mux.Vars(req)
	path := appPath + params["name"] + "miniature.img"
	http.FileServer(http.Dir(path))
}
