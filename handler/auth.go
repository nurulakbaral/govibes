package handler

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"govibes.app/config"
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
	if err := userEntity.SelectRowByEmail(c.Context(), reqBody.Email); err != nil {
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

	jwtToken := jwt.New(jwt.SigningMethodHS256)
	jwtClaims := jwtToken.Claims.(jwt.MapClaims)
	jwtClaims["user_id"] = userEntity.Id
	jwtClaims["email"] = userEntity.Email
	jwtClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	jwtSigned, err := jwtToken.SignedString([]byte(config.Config(config.JWT_SECRET_KEY)))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Generate jwt token is error",
			"errors":  err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "Access-Token",
		Value:    jwtSigned,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"status":      "success",
		"message":     "Login successfully!",
		"accessToken": jwtSigned,
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
