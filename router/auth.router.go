package router

import (
	"main/controller"
	"main/database"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(api fiber.Router, db *database.Database){
	authHandler := controller.AuthHandler(db);
	api.Get("/login", authHandler.UserLogin)
	api.Post("/reg", authHandler.UserRegistration)
}