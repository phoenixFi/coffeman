package handlers

import (
	"encoding/json"
	"net/http"
	"myapp/database"
	"myapp/models"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var cat models.Category
	json.NewDecoder(r.Body).Decode(&cat)
	database.DB.Create(&cat)
	json.NewEncoder(w).Encode(cat)
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	var cats []models.Category
	database.DB.Find(&cats)
	json.NewEncoder(w).Encode(cats)
}
