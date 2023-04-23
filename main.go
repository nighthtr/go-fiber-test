// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nighthtr/go-fiber-test/database"
	"github.com/nighthtr/go-fiber-test/routes"
)

func init() {
	// Чтение переменных из .env
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Подключение к базе данных
	database.ConnectDB()
}

func main() {
	// Создание экземпляра Fiber
	app := fiber.New()

	// Регистрация маршрутов
	routes.SetupRoutes(app)

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("Server is running on http://localhost:" + port)
	log.Fatal(app.Listen(":" + port))
}
