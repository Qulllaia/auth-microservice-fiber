package controller

import (
	"main/database"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
    db *database.Database
}

func UserHandler(db *database.Database) *UserController {
    return &UserController{db: db}
}

func (uc *UserController) GetUsers(c *fiber.Ctx) error{

	if user, err:=uc.db.GetUsers(); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
		}else{
			return c.Status(fiber.StatusOK).JSON(user);
	}
}

func (uc *UserController) PostUser(c *fiber.Ctx) error{

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
			"ID": user.ID,
			"login": user.Email,
		})
	}
}

func (uc *UserController) GetUserWithId(c *fiber.Ctx) error{
	params, err := c.ParamsInt("ID");

	if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

	if user, err := uc.db.GetUserWithId(params); err != nil{

		return c.SendStatus(fiber.StatusInternalServerError);

	}else{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ID": user.ID,
			"email": user.Email,
		})
	}

}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error{
	param, err := c.ParamsInt("ID");

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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid request body",
		});
	}

	if err := uc.db.PutUser(user); err != nil{
		return c.SendStatus(fiber.StatusInternalServerError);
	}else{
		return c.Status(fiber.StatusOK).JSON(user);
	}
	
}