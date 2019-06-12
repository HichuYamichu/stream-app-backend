package app

import (
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// ServeMiniature : servs video miniature
func ServeMiniature(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	// G:\Videos\stream-app-vids\fuk
	path := "G:/Videos/stream-app-vids/" + params["id"] + "/" + params["id"] + ".jpeg"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	jpeg.Encode(res, img, nil) // Write to the ResponseWriter
}
