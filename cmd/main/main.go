package main

import (
	"fmt"
	"net/http"

	"pdrygala.com/go-bookstore-api/pkg/routes"
)

func main() {
	// Create the REST API server
	apiHandler := routes.NewServer()

	if err := http.ListenAndServe(":8080", apiHandler); err != nil {
		fmt.Println("HTTP server error:", err)
	}

}
