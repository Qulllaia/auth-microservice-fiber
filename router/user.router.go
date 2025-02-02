package router

import (
	"main/controller"

	"github.com/gofiber/fiber/v2"
)
func Router(api fiber.Router){
	api.Get("/", controller.GetUser)
	api.Post("/", controller.PostUser)
	api.Get("/:id", controller.GetUserWithId)
	api.Put("/:id", controller.PutUser)
	api.Delete("/:id", controller.DeleteUser)
}