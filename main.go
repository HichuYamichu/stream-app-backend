package main

import (
	"fmt"
	router "github.com/HichuYamichu/stream-app-server/router"
	MongoDB "github.com/HichuYamichu/stream-app-server/storage"
	"github.com/gorilla/handlers"
	"net/http"
)

func init() {
	MongoDB.Connect()
	if MongoDB.DB != nil {
		fmt.Println("Connected to MongoDB!")
	}
}

func main() {
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	appRouter := router.Get()
	fmt.Println("Serving on port 3000")
	http.ListenAndServe("127.0.0.1:3000", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(appRouter))
}
