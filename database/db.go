package database

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type User struct {
    gorm.Model
    Name  string
    Email string
}

var DB *gorm.DB

func Connect() {
	dsn := "user=postgres password=15456910 dbname=Coffeman sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Ошибка подключения:", err)
	}

	fmt.Println("✅ Успешное подключение к базе данных!")

	db.Migrator().CreateTable(&User{})
	DB = db
}
