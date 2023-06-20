package main

import (
	"github.com/gemstack/jobs-api/routes"
)

func main() {
	//r := gin.Default()

	// Define your routes and handlers here
	r := routes.SetupRoutes()

	r.Run(":8080")
}
