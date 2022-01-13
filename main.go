package main

import (
	"golang/database"
	"golang/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to an Awesome API")
}

func setupRoutes(app *fiber.App) {
	// Welcome endpoint
	app.Get("/api", welcome)
	// User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Post("/api/users/getuser", routes.GetUser)
	app.Post("/api/users/update", routes.UpdateUser)
	app.Post("/api/users/delete", routes.DeleteUser)
}

func main() {
	database.ConnectDb()

	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
