package router

import (
	"kasir-api/handler"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/api/product", handler.ProdukHandler)
	http.HandleFunc("/api/product/", handler.ProdukByIDHandler)

	http.HandleFunc("/categories", handler.CategoryHandler)
	http.HandleFunc("/categories/", handler.CategoryByIDHandler)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"OK"}`))
	})
}
