package router

import (
	"main/controller"
	"main/database"
	"main/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router, db *database.Database){
	userHandler := controller.UserHandler(db);
	api.Get("/",  middleware.AuthMiddleware, userHandler.GetUsers)
	api.Post("/", middleware.AuthMiddleware, userHandler.PostUser)
	api.Get("/:id", middleware.AuthMiddleware, userHandler.GetUserWithId)
	api.Put("/", middleware.AuthMiddleware, userHandler.PutUser)
	api.Delete("/:id", middleware.AuthMiddleware, userHandler.DeleteUser)
}