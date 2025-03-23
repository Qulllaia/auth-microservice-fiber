package router

import (
	"main/controller"
	"main/database"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(api fiber.Router, db *database.Database){
	subApi := api.Group("/auth")
	authHandler := controller.AuthHandler(db);
	subApi.Get("/login", authHandler.UserLogin)
	subApi.Post("/reg", authHandler.UserRegistration)
}