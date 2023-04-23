// middlewares/middlewares.go
package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nighthtr/go-fiber-test/models"
)

// AuthMiddleware проверяет авторизацию пользователя
func AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Здесь можно выполнить проверку авторизации пользователя, например, проверить наличие токена в заголовках запроса,
		// и/или проверить его на валидность в базе данных

		// В данном примере пропускаем все запросы без проверки авторизации
		return ctx.Next()
	}
}

// ErrorHandlerMiddleware обрабатывает ошибки и возвращает соответствующий HTTP-ответ
func ErrorHandlerMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()

		// Обработка ошибок
		if err != nil {
			// Если ошибка - "404 Not Found"
			if err == fiber.ErrNotFound {
				return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Not Found",
				})
			}

			// Если ошибка - "500 Internal Server Error"
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}

		return nil
	}
}

// SetCurrentUserMiddleware устанавливает текущего пользователя в контексте запроса
func SetCurrentUserMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Здесь можно выполнить логику установки текущего пользователя в контексте запроса,
		// например, получить пользователя из базы данных по токену авторизации

		// В данном примере создаем пустого пользователя и устанавливаем его в контексте
		user := models.User{}
		ctx.Locals("user", &user)
		return ctx.Next()
	}
}
