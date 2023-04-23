// controllers/user_controller.go
package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nighthtr/go-fiber-test/database"
	"github.com/nighthtr/go-fiber-test/models"
)

// UserController контроллер пользователей
type UserController struct {
}

// NewUserController создает новый экземпляр контроллера пользователей
func NewUserController() *UserController {
	return &UserController{}
}

// Create обрабатывает запрос на создание нового пользователя
func (c *UserController) Create(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return ctx.JSON(user)
}

// GetAll обрабатывает запрос на получение списка пользователей
func (c *UserController) GetAll(ctx *fiber.Ctx) error {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get users",
		})
	}

	return ctx.JSON(users)
}

// GetByID обрабатывает запрос на получение информации о пользователе по ID
func (c *UserController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return ctx.JSON(user)
}

// UpdateByID обрабатывает запрос на обновление информации о пользователе по ID
func (c *UserController) UpdateByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result = database.DB.Save(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	return ctx.JSON(user)
}

// DeleteByID обрабатывает запрос на удаление пользователя по ID
func (c *UserController) DeleteByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.User
	result := database.DB.Delete(&user, id)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
