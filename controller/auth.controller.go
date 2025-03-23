package controller

import (
	"main/database"
	"main/database/dto"
	"main/models"
	"os"
	"strconv"
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
			"User registration error": err,
		})
	}

	signedToken, err := tokenGenerator(user);
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"The token generation process ended with an error:": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": signedToken,
	}) 
}

func (uc *AuthController) UserLogin(c *fiber.Ctx) error{

	var loginData dto.UserLoginDto;
	var resultUser models.UserData;

	if err := c.BodyParser(&loginData); err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if resultUser, err := uc.db.LoginUser(loginData); err != nil || resultUser.ID == 0{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Wrong Credentials",
		})
	}

	if resultToken, err := tokenGenerator(resultUser); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "The token generation process ended with an error",
			"message": err,
		});
	}else{
		
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": resultToken,
		});
	}

}

func tokenGenerator(user models.UserData) (string, error){
	claims := dto.AuthDto{
		ID: user.ID, 
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),      
			ID:        strconv.Itoa(user.ID),   
			Audience:  []string{"client-app"}, 
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims);
	return token.SignedString(jwtKey);
}
