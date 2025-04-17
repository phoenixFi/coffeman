package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		json.NewEncoder(w).Encode(Response{"fail", "Только POST-запросы"})
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Message == "" {
		json.NewEncoder(w).Encode(Response{"fail", "Некорректное JSON-сообщение"})
		return
	}

	fmt.Println("Получено сообщение:", req.Message)
	json.NewEncoder(w).Encode(Response{"success", "Данные успешно приняты"})
}

func madin(){
	http.HandleFunc("/", handler)
	fmt.Println("Сервер работает на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
