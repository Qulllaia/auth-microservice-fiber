package controller

import "github.com/gofiber/fiber/v2"

func GetUser(c *fiber.Ctx) error{
	return c.SendString("Hello Get")
}

func PostUser(c *fiber.Ctx) error{
	return c.SendString("Hello Post")
}

func GetUserWithId(c *fiber.Ctx) error{
	params := c.AllParams()
	return c.SendString("Hello Get With Id: " + params["id"])
}

func DeleteUser(c *fiber.Ctx) error{
	params := c.AllParams()
	return c.SendString("Hello Delete With Id: " + params["id"])
}

func PutUser(c *fiber.Ctx) error{
	params := c.AllParams()
	return c.SendString("Hello Put With Id: " + params["id"])
}