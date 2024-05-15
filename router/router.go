package router

import (
	"govibes.app/handler"
	"govibes.app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/user")
	user.Post("/register", handler.Register)
	user.Get("/all", middleware.Protected(), handler.GetAllUser)
	user.Put("/edit-profile", middleware.Protected(), handler.EditProfile)
}
