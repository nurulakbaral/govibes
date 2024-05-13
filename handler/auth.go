package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Row struct {
	Id       int
	Username string
	Email     string
}

func Login(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status": "success",
		"message": "Login successfully!",
	})
}