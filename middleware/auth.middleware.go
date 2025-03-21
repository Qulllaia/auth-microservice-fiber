package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("SECRET_TOKEN"));

func AuthMiddleware(c *fiber.Ctx) error{

	tokenString := strings.TrimSpace(
		strings.TrimPrefix(
			c.Get("Authorization"), "Bearer"));

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