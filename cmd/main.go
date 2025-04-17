package main

import (
	"log"
	"net/http"
	"os"

	"myapp/database"
	"myapp/router"
)

func main() {
	database.Connect()
	r := router.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Println("Inventory service running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
