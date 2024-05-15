package handler

import (
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

	reqBody := new(model.RequestUserRegister)

	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Parser is error: Please, review your input",
			"errors":  err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body: Please, review yout input",
			"errors":  err.Error(),
		})
	}

	hashedPassword, err := utils.HashPassword(reqBody.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Hash password is error: Please, check your input",
			"errors":  err.Error(),
		})
	}
	reqBody.Password = hashedPassword

	userEntity := new(model.User)
	if err := userEntity.InsertUser(c.Context(), *reqBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create user: Unable to insert new user data to database",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Created new user successfully",
		"data": model.ResponseUserRegister{
			Id:        userEntity.Id,
			Name:      userEntity.Name,
			Username:  userEntity.Username,
			Email:     userEntity.Email,
			CreatedAt: userEntity.CreatedAt,
			DeletedAt: userEntity.DeletedAt,
		},
	})
}
