package main

import (
	"fmt"
	"net/http"

	"kasir-api/router"
)

func main() {
	router.RegisterRoutes()

	fmt.Println("Server running di :8080")
	http.ListenAndServe(":8080", nil)
}
