package main

import "github.com/gofiber/fiber/v2"
import "github.com/bengin34/depixen/database"

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)


	app.Listen(":3000")
}

