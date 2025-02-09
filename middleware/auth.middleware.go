package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret-key");

func AuthMiddleware(c *fiber.Ctx) error{

	tokenString := c.Get("Authorization");

	if tokenString == ""{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"Unauthorized",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		return jwtKey, nil
	})

	if err != nil || !token.Valid{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"Invalid Token",
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user", claims)

	return c.Next();
}