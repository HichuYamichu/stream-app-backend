package videos

import (
	"context"
	"fmt"
	"net/http"
)

const appPath = "G:/Videos/streamer-app-vids/"

var ctx = context.TODO()

func handleErr(err error, res http.ResponseWriter) {
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
}
