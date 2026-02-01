package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"kasir-api/util"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// HandleCategorys - GET /api/category
func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	Categorys, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("hit endpoint /Category gagal: %T %v", err, err)
		return
	}

	util.JSON(w, http.StatusCreated, Categorys)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var Category models.Category
	err := json.NewDecoder(r.Body).Decode(&Category)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.Create(&Category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.JSON(w, http.StatusCreated, Category)
}

// HandleCategoryByID - GET/PUT/DELETE /api/category/{id}
func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetByID - GET /api/category/{id}
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	Category, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	util.JSON(w, http.StatusCreated, Category)

}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var Category models.Category
	err = json.NewDecoder(r.Body).Decode(&Category)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	Category.ID = id
	err = h.service.Update(&Category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.JSON(w, http.StatusCreated, Category)
}

// Delete - DELETE /api/category/{id}
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.JSON(w, http.StatusOK, map[string]string{"message": "Category deleted successfully"})

}
