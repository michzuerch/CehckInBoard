package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/michzuerch/CheckInBoard/database"
	"github.com/michzuerch/CheckInBoard/models"
)

type User struct {
	//This is not the model user, see this as the serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {

	return User{
		ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)

	responseUsers := []User{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)

	}

	return c.Status(fiber.StatusOK).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(user, "id = ?", id)

	if user.ID == 0 {
		return errors.New("User not found")
	}

	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(":id must be integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(fiber.StatusOK).JSON(responseUser)
}
