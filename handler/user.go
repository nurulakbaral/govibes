package handler

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"govibes.app/model"
	"govibes.app/utils"
)

func GetAllUser(c *fiber.Ctx) error {
	user := new(model.User)
	rows, err := user.SelectAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Unable to get all user data",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Get all data successfully",
		"data":    rows,
	})
}

func Register(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Parser is error: Please, review your input",
			"errors":  err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(new(model.RequestUser)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body: Please, review yout input",
			"errors":  err.Error(),
		})
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Hash password is error: Please, check your input",
			"errors":  err.Error(),
		})
	}

	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	if err := user.InsertUser(c.Context()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create user: Unable to insert new user data to database",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Created new user successfully",
		"data": model.ResponseUser{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			DeletedAt: user.DeletedAt,
		},
	})
}
