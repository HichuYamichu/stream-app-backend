package main

import (
	"fmt"
	"net/http"

	"github.com/HichuYamichu/stream-app-server/api"
	"github.com/HichuYamichu/stream-app-server/db"
)

func main() {
	db.Connect()
	appRouter := api.GetRouter()
	fmt.Println("Serving on port 3000")
	http.ListenAndServe("127.0.0.1:3000", appRouter)
}
