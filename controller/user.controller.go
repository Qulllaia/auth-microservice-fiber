package controller

import (
	"main/database"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
    db *database.Database
}

func NewUserController(db *database.Database) *UserController {
    return &UserController{db: db}
}

func (uc *UserController) GetUser(c *fiber.Ctx) error{
	return c.SendString("Hello Get")
}

func (uc *UserController) PostUser(c *fiber.Ctx) error{

	var user models.UserData;

	if err := c.BodyParser(&user); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid request body",
		});
	}
	
	if err := uc.db.CreateUser(&user); err != nil{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ID": user.User_id,
			"login": user.Login,
		})
	}

	return c.SendStatus(fiber.StatusInternalServerError);
}

func (uc *UserController) GetUserWithId(c *fiber.Ctx) error{
	params, err := c.ParamsInt("ID");

	if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

	if user, err := uc.db.GetUser(params); err == nil{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ID": user.User_id,
			"login": user.Login,
		})
	}

	return c.SendStatus(fiber.StatusInternalServerError);
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error{
	params := c.AllParams()
	return c.SendString("Hello Delete With Id: " + params["id"])
}

func (uc *UserController) PutUser(c *fiber.Ctx) error{
	params := c.AllParams()
	return c.SendString("Hello Put With Id: " + params["id"])
}