package main

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"myapp/database"
	"myapp/models"
)

func main() {
	// Подключение к БД
	database.Connect()

	// Создание роутера
	r := mux.NewRouter()

	// API маршруты
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/users", GetAllUsers).Methods("GET")

	// Раздача статических файлов
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(fs)

	// Запуск сервера
	log.Println("🚀 Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// 🔹 Добавление пользователя
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

// 🔹 Получение пользователя по ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user models.User
	if err := database.DB.First(&user, vars["id"]).Error; err != nil {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// 🔹 Обновление пользователя
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user models.User
	if err := database.DB.First(&user, vars["id"]).Error; err != nil {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

// 🔹 Удаление пользователя
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user models.User
	if err := database.DB.Delete(&user, vars["id"]).Error; err != nil {
		http.Error(w, "Ошибка удаления", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// 🔹 Получение всех пользователей
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}
