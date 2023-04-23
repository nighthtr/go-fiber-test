// database/database.go
package database

import (
	"fmt"
	"log"
	"os"

	"github.com/nighthtr/go-fiber-test/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB осуществляет подключение к базе данных PostgreSQL и возвращает ссылку на экземпляр DB
func ConnectDB() {
	// Чтение данных для подключения из переменных окружения
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Формирование строки подключения к базе данных
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	// Подключение к базе данных
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Connected to database")

	// Миграции модели User
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to auto migrate User model: %v", err)
	}
}
