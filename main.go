package main

import (
	"netflix-gpt-backend/config"
	"netflix-gpt-backend/router"
)

func main() {

	config.ConnectDatabase()

	// Initialize the router from the router package
	r := router.SetupRouter()

	// Start the server
	r.Run(":8080")
}
