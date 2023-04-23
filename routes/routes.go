// routes/routes.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nighthtr/go-fiber-test/controllers"
	"github.com/nighthtr/go-fiber-test/middlewares"
)

// SetupRoutes устанавливает маршруты для API с пользователями
func SetupRoutes(app *fiber.App) {
	// Инициализация контроллера
	userController := controllers.NewUserController()

	// Группа маршрутов для API с пользователями
	users := app.Group("/api/users")

	// Защита маршрутов с использованием middleware
	users.Use(middlewares.AuthMiddleware())

	// Маршрут для создания пользователя
	users.Post("/", userController.Create)

	// Маршрут для получения списка пользователей
	users.Get("/", userController.GetAll)

	// Маршрут для получения информации о пользователе по ID
	users.Get("/:id", userController.GetByID)

	// Маршрут для обновления информации о пользователе по ID
	users.Put("/:id", userController.UpdateByID)

	// Маршрут для удаления пользователя по ID
	users.Delete("/:id", userController.DeleteByID)
}
