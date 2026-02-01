package main

import (
	"kasir-api/config"
	"kasir-api/router"
	"log"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	cfg := config.Load()
	
	log.Printf("Port: %s", cfg.Port)
	log.Printf("DB Connection: %s", cfg.DBConn)
	
	if cfg.Port == "" {
		log.Fatal("PORT is empty")
	}
	if cfg.DBConn == "" {
		log.Fatal("DB_CONN is empty")
	}
	
	// register routes SEBELUM server start
	db, err := config.InitDB(cfg.DBConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router.RegisterRoutes(db)

	addr := "0.0.0.0:" + cfg.Port
	log.Println("APP START")
	log.Printf("Startup time: %v\n", time.Since(start))
	log.Println("Server running di", addr)

	// start server (sekali saja)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
