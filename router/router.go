package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"myapp/handlers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Product routes
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", handlers.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PATCH")
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/products", handlers.ListProducts).Methods("GET")

	// Category routes
	r.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	r.HandleFunc("/categories", handlers.GetAllCategories).Methods("GET")

	// Static (optional)
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(fs)

	return r
}
