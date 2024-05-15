package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"govibes.app/model/user"
	"govibes.app/utils"
)

func Login(c *fiber.Ctx) error {

	reqBody := new(user.RequestLogin)

	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "login is failed: review your input",
			"errors":  err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "login is failed: your data type is wrong",
			"errors":  err.Error(),
		})
	}

	userEntity := new(user.User)
	if err := userEntity.SelectByEmail(c.Context(), reqBody.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "User is not found",
			"errors":  err.Error(),
		})
	}

	if err := utils.ValidateHashedPassword(reqBody.Password, userEntity.Password); err != nil {
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
		"data": user.ResponseLogin{
			Id:        userEntity.Id,
			Name:      userEntity.Name,
			Username:  userEntity.Username,
			Email:     userEntity.Email,
			CreatedAt: userEntity.CreatedAt,
			DeletedAt: userEntity.DeletedAt,
		},
	})
}
