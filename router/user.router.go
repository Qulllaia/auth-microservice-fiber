package router

import (
	"main/controller"
	"main/database"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router, db *database.Database){
	userController := controller.NewUserController(db);
	api.Get("/", userController.GetUser)
	api.Post("/", userController.PostUser)
	api.Get("/:id", userController.GetUserWithId)
	api.Put("/:id", userController.PutUser)
	api.Delete("/:id", userController.DeleteUser)
}