package main

import (
	"main/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Route("/api", router.Router)
	app.Listen(":3000")
}