package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"kasir-api/model"
	"kasir-api/store"
	"kasir-api/util"
)

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		util.JSON(w, http.StatusOK, store.CategoryData)

	case http.MethodPost:
		var c model.Category
		json.NewDecoder(r.Body).Decode(&c)
		c.ID = len(store.CategoryData) + 1
		store.CategoryData = append(store.CategoryData, c)
		util.JSON(w, http.StatusCreated, c)
	}
}

func CategoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.Error(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	for i, c := range store.CategoryData {
		if c.ID == id {
			switch r.Method {
			case http.MethodGet:
				util.JSON(w, http.StatusOK, c)
				return

			case http.MethodPut:
				var update model.Category
				json.NewDecoder(r.Body).Decode(&update)
				update.ID = id
				store.CategoryData[i] = update
				util.JSON(w, http.StatusOK, update)
				return

			case http.MethodDelete:
				store.CategoryData = append(store.CategoryData[:i], store.CategoryData[i+1:]...)
				util.JSON(w, http.StatusOK, map[string]string{"message": "deleted"})
				return
			}
		}
	}

	util.Error(w, http.StatusNotFound, "Category tidak ditemukan")
}
