package routes

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/michzuerch/CheckInBoard/database"
	"github.com/michzuerch/CheckInBoard/restmodels"
)

type Stamp struct {
	//This is not the model user, see this as the serializer
	ID      uint     `json:"id"`
	CheckIn bool     `json:"checkin"`
	Stamp   datetime `json:"stamp"`
}

func CreateResponseStamp(stampModel restmodels.Stamp) Stamp {

	return Stamp{
		ID: stampModel.ID, CheckIn: stampModel.CheckIn, Stamp: stampModel.Stamp,
	}
}

func CreateStamp(c *fiber.Ctx) error {
	var user restmodels.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func GetStamp(c *fiber.Ctx) error {
	users := []restmodels.User{}
	database.Database.Db.Find(&users)

	responseUsers := []User{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)

	}

	return c.Status(fiber.StatusOK).JSON(responseUsers)
}

func findStamp(id int, user *restmodels.User) error {
	database.Database.Db.Find(user, "id = ?", id)

	if user.ID == 0 {
		return errors.New("User not found")
	}

	return nil
}

func GetStamp(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user restmodels.User

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(":id must be integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func UpdateStamp(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	fmt.Printf("UpdateUser(%v)", id)

	var user restmodels.User

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(":id must be integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)

}

func DeleteStamp(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	fmt.Printf("DeleteUser(%v)", id)

	var user restmodels.User

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(":id must be integer")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Successfully deleted user")
}
