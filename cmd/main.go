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
		Addr:         "0.0.0.0:3001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Serving on port 3001")
	log.Fatal(srv.ListenAndServe())
}
