package router

import (
	"main/controller"
	"main/database"
	"main/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router, db *database.Database){
	subApi := api.Group("/users")
	userHandler := controller.UserHandler(db);
	subApi.Get("/",  middleware.AuthMiddleware, userHandler.GetUsers)
	subApi.Post("/", middleware.AuthMiddleware, userHandler.PostUser)
	subApi.Get("/:id", middleware.AuthMiddleware, userHandler.GetUserWithId)
	subApi.Put("/", middleware.AuthMiddleware, userHandler.PutUser)
	subApi.Delete("/:id", middleware.AuthMiddleware, userHandler.DeleteUser)
}