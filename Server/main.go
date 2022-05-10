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
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Post("/api/users", routes.CreateUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

}

func main() {
	database.ConnectDb()

	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
