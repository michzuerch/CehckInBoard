package main

//go:generate sqlboiler --wipe psql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/michzuerch/CheckInBoard/routes"
)

// "github.com/michzuerch/CheckInBoard/database"

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

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("postgres", `dbname=CheckInBoard host=localhost user=michzuerch password=lx0lc33a sslmode=disable`)

	dieIf(err)

	err = db.Ping()
	dieIf(err)

	fmt.Println("Database connected")

	// database.ConnectDb()

	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
