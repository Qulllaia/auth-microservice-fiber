package controller

import (
	"fmt"
	"main/database"
	"main/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
    db *database.Database
}

func UserHandler(db *database.Database) *UserController {
    return &UserController{db: db}
}

func (uc *UserController) GetUsers(c *fiber.Ctx) error{

	if user, err:=uc.db.GetUsers(); err == nil{
		return c.Status(fiber.StatusOK).JSON(user);
	}

	return c.SendStatus(fiber.StatusInternalServerError);
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

	if user, err := uc.db.GetUserWithId(params); err == nil{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ID": user.User_id,
			"login": user.Login,
		})
	}

	return c.SendStatus(fiber.StatusInternalServerError);
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error{
	param, err := strconv.Atoi(c.Params("id"))

	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
	}

	if err:= uc.db.DeleteUser(param); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
	}else{
		c.SendStatus(fiber.StatusOK);
	}

	return c.SendStatus(fiber.StatusInternalServerError);
}

func (uc *UserController) PutUser(c *fiber.Ctx) error{
	var user models.UserData;

	if err:= c.BodyParser(&user); err != nil{
		fmt.Println(user.Login);
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid request body",
		});
	}

	if err := uc.db.PutUser(user); err == nil{
		return c.Status(fiber.StatusOK).JSON(user);
	}
	
	return c.SendStatus(fiber.StatusInternalServerError);
}