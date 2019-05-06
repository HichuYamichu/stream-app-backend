package dispatcher

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func SendVideo(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	location := "g:/Videos/stream-app-vids/" + params["name"] + "/" + params["name"] + ".mp4"
	fmt.Println(location)
	http.ServeFile(res, req, location)
}
