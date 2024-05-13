package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Login successfully!",
	})
}
