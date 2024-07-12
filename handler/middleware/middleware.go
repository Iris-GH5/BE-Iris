package middleware

import (
	"github.com/Iris-GH5/BE-Iris/database"
	"github.com/Iris-GH5/BE-Iris/model/entity"
	"github.com/Iris-GH5/BE-Iris/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token format",
		})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	validToken, err := utils.VerifyToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var existingUser entity.User
	err = database.DB.Where("email = ?", validToken["email"]).First(&existingUser).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return ctx.Next()
}