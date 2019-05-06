package videos

import (
	"fmt"
	"github.com/gorilla/mux"
	"image/jpeg"
	"log"
	"net/http"
	"os"
)

func ServeVideoMiniatures(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	// G:\Videos\stream-app-vids\fuk
	path := "G:/Videos/stream-app-vids/" + params["name"] + "/" + params["name"] + ".jpeg"
	fmt.Println(path)
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
