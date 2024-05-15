package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"govibes.app/model/user"
	"govibes.app/utils"
)

func GetAllUser(c *fiber.Ctx) error {
	user := new(user.User)
	rows, err := user.SelectRows(c.Context())

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
	reqBody := new(user.RequestRegister)

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

	userEntity := new(user.User)
	if err := userEntity.InsertRow(c.Context(), *reqBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create user: Unable to insert new user data to database",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Created new user successfully",
		"data": user.ResponseRegister{
			Id:        userEntity.Id,
			Name:      userEntity.Name,
			Username:  userEntity.Username,
			Email:     userEntity.Email,
			CreatedAt: userEntity.CreatedAt,
			DeletedAt: userEntity.DeletedAt,
		},
	})
}

func EditProfile(c *fiber.Ctx) error {
	reqBody := new(user.RequestEditProfile)

	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Parsing request body is error",
			"errors":  err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Validating request body is error",
			"errors":  err.Error(),
		})
	}
	userStore := c.Locals("user").(*jwt.Token)
	claims := userStore.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	reqBody.Id = uuid.MustParse(userId)

	userEntity := new(user.User)
	if err := userEntity.UpdateRowName(c.Context(), *reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Edit profile is failed, try again",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "success",
		"message": "Edit profile is success",
		"data": user.ResponseEditProfile{
			Id:   userEntity.Id,
			Name: userEntity.Name,
		},
	})
}
