package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"myapp/database"
	"myapp/models"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.DB.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		http.Error(w, "Продукт не найден", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		http.Error(w, "Продукт не найден", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&product)
	database.DB.Save(&product)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		http.Error(w, "Ошибка при удалении", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	// Фильтрация и пагинация
	category := r.URL.Query().Get("category")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	tx := database.DB

	if category != "" {
		tx = tx.Where("category_id = ?", category)
	}
	if limitStr != "" {
		limit, _ := strconv.Atoi(limitStr)
		tx = tx.Limit(limit)
	}
	if offsetStr != "" {
		offset, _ := strconv.Atoi(offsetStr)
		tx = tx.Offset(offset)
	}

	tx.Find(&products)
	json.NewEncoder(w).Encode(products)
}
