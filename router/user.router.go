package router

import (
	"main/controller"
	"main/database"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router, db *database.Database){
	userHandler := controller.UserHandler(db);
	api.Get("/", userHandler.GetUsers)
	api.Post("/", userHandler.PostUser)
	api.Get("/:id", userHandler.GetUserWithId)
	api.Put("/", userHandler.PutUser)
	api.Delete("/:id", userHandler.DeleteUser)
}