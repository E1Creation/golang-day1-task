package router

import (
	"database/sql"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	// Setup routes
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	http.HandleFunc("/api/product", productHandler.HandleProducts)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)

	http.HandleFunc("/api/category", categoryHandler.HandleCategories)
	http.HandleFunc("/api/category/", categoryHandler.HandleCategoryByID)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"OK"}`))
	})
}
