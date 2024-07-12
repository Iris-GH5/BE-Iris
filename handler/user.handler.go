package handler

import (
	"github.com/Iris-GH5/BE-Iris/database"
	"github.com/Iris-GH5/BE-Iris/model/dto"
	"github.com/Iris-GH5/BE-Iris/model/entity"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(c *fiber.Ctx) error {
	user := new(dto.UserRegisterRequestDTO)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing new user",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error validating new user",
			"error":   errValidate.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error hashing password",
			"error":   err.Error(),
		})
	}

	newUser := entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  string(hashedPassword),
	}
	
	var existingUser entity.User
	res := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if res.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	newUserRes := database.DB.Create(&newUser)

	if newUserRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new user",
			"error":   newUserRes.Error.Error(),
		})
	}

	responseDTO := dto.UserRegisterResponseDTO{
		Message:   "New user created successfully",
		ID:        newUser.ID,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
	}

	return c.Status(201).JSON(responseDTO)
}