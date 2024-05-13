package main

import (
	"log"

	"govibes.app/database"
	"govibes.app/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Govibes",
		AppName:       "Govibes",
	})

	database.Connect()
	router.Init(app)

	log.Fatal(app.Listen(":8001"))
}
