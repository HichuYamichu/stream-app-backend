package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/HichuYamichu/stream-app-server/db"
	"github.com/HichuYamichu/stream-app-server/router"
	"github.com/gorilla/handlers"
)

func main() {
	db.Connect()
	r := router.GetRouter()

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, r),
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Serving on port 3000")
	log.Fatal(srv.ListenAndServe())
}
