package controller

import (
	"main/database"
	"main/database/dto"
	"main/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
    db *database.Database
}

var jwtKey = []byte(os.Getenv("SECRET_TOKEN"));

func AuthHandler(db *database.Database) *AuthController {
    return &AuthController{db: db}
}

func (uc *AuthController) UserRegistration(c *fiber.Ctx) error{

	var user models.UserData;

	if err := c.BodyParser(&user); err !=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid request body",
		});
	}
	
	if err := uc.db.CreateUser(&user); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Ошибка при регистрации пользователя": err,
		})
	}

	claims := dto.AuthDto{
		ID: user.ID, 
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),      
			Issuer:    "my-app",                    
			Subject:   "user-auth",
			ID:        "unique-id",   
			Audience:  []string{"client-app"}, 
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Ошибка при создании токена:": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": signedToken,
	}) 
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
