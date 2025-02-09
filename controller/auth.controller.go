package controller

import (
	"main/database"
	"main/database/dto"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
    db *database.Database
}

func AuthHandler(db *database.Database) *AuthController {
    return &AuthController{db: db}
}

func (uc *AuthController) UserRegistration(c *fiber.Ctx) error{

	var user models.UserData;

	if err := c.BodyParser(&user); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid request body",
		});
	}
	
	if err := uc.db.CreateUser(&user); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}else{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ID": user.User_id,
			"login": user.Login,
		})
	}
}

func (uc *AuthController) UserLogin(c *fiber.Ctx) error{

	var loginData dto.UserLoginDto

	if err := c.BodyParser(&loginData); err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if result, err := uc.db.LoginUser(loginData); result || err == nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Wrong Credentials",
		})
	}else{
		return c.SendStatus(200);
	}

}
