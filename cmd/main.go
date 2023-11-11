package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/niyioo/backend/pkg/database"
	"github.com/niyioo/backend/pkg/routes"
)

func main() {
	database.ConnectDB() // Connect to MongoDB

	// Set user and post routes
	routes.SetUserRoutes()
	routes.SetPostRoutes()

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
