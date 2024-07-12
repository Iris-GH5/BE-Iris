package route

import (
	"github.com/Iris-GH5/BE-Iris/handler"
	"github.com/Iris-GH5/BE-Iris/handler/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())

	api := app.Group("/api")

	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.Get("/current", middleware.Auth, handler.GetCurrentUser)
	users.Post("/login", handler.UserLogin)
	users.Post("/talent/register", handler.UserRegister)
}
