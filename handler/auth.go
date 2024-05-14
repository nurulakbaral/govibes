package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"govibes.app/model"
	"govibes.app/utils"
)

func Login(c *fiber.Ctx) error {
	type RequestUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	requestUser := new(RequestUser)

	if err := c.BodyParser(requestUser); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "login is failed: review your input",
			"errors":  err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(requestUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "login is failed: your data type is wrong",
			"errors":  err.Error(),
		})
	}

	userEntity := new(model.User)
	if err := userEntity.SelectByEmail(c.Context(), requestUser.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "User is not found",
			"errors":  err.Error(),
		})
	}

	if err := utils.ValidateHashedPassword(requestUser.Password, userEntity.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "User is not found",
			"errors":  err.Error(),
		})
	}

	// Todo: Create JWT & Set Cookie

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Login successfully!",
		"data":    userEntity,
	})
}
