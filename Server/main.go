package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/michzuerch/CheckInBoard/database"
	"github.com/michzuerch/CheckInBoard/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("API CheckInBoard V1")
}

func setupRoutes(app *fiber.App) {
	// welcome endpoint
	app.Get("/api", welcome)

	// user endpoints
	app.Post("/api/users", routes.CreateUser)

}

func main() {
	database.ConnectDb()

	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
