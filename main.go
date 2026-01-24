package main

import (
	"fmt"
	"net/http"

	"kasir-api/router"
	"log"
	"time"
)

func main() {
	start := time.Now()

	log.Println("APP START")
	log.Printf("Startup time: %v\n", time.Since(start))
	fmt.Println("Server running di :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
	router.RegisterRoutes()

	http.ListenAndServe(":8080", nil)
}
