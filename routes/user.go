package routes

import (
	"fmt"
	"golang/database"
	"golang/models"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	// This is not the model, more like a serializer
	Username        string `json:"username"`
	Password        string `json:"password"`
	Confirmpassword string `json:"confirmpassword"`
}

func CreateResponseUser(user models.User) User {
	return User{Username: user.Username, Password: user.Password, Confirmpassword: user.Confirmpassword}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if len(user.Username) < 6 {
		return c.Status(400).JSON("username too short")
	}

	if len(user.Password) < 6 {
		return c.Status(400).JSON("password too short")
	}

	if user.Confirmpassword != user.Password {
		return c.Status(400).JSON("password not match")
	}
	var findUser models.User
	database.Database.Db.Find(&findUser, "username= ?", user.Username)
	if findUser.Username != user.Username {
		database.Database.Db.Create(&user)
		responseUser := CreateResponseUser(user)
		return c.Status(200).JSON(responseUser)
	} else {
		return c.Status(400).JSON("username ton tai")
	}

}
func GetUser(c *fiber.Ctx) error {
	var user models.User

	type LogUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var logUser LogUser

	if err := c.BodyParser(&logUser); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	database.Database.Db.Find(&user, "username= ?", logUser.Username)
	if logUser.Username == user.Username && logUser.Password == user.Password {
		return c.Status(400).JSON("login successfuly")

	}
	return c.Status(400).JSON("login not successfuly")

}

func UpdateUser(c *fiber.Ctx) error {

	var user models.User

	type UpdateUser struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		NewPassword     string `json:"newpassword"`
		Confirmpassword string `json:"confirmpassword"`
	}
	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	database.Database.Db.Find(&user, "username= ?", updateData.Username)
	if updateData.Username == user.Username && updateData.Password == user.Password {
		if updateData.Confirmpassword == updateData.NewPassword {
			user.Password = updateData.NewPassword
			user.Confirmpassword = updateData.NewPassword
			database.Database.Db.Save(&user)

			responseUser := CreateResponseUser(user)

			return c.Status(200).JSON(responseUser)
		}
		return c.Status(400).JSON("password not match")
	}

	fmt.Print(updateData)
	return c.Status(400).JSON("username not right & pass not right")
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User

	type dlUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var dlus dlUser

	if err := c.BodyParser(&dlus); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	database.Database.Db.Find(&user, "username= ?", user.Username)

	if dlus.Username == user.Username && dlus.Password == user.Password {
		database.Database.Db.Delete(&user)
		return c.Status(200).JSON("Successfully deleted User")
	}

	return c.Status(400).JSON("username not right & password not right")

}
