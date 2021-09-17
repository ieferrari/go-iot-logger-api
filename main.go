package main

import (
	"go-iot-logger-api/database"
	"go-iot-logger-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load() // load .env file in this directory, then use os.Getenv("PSWD")
	app := fiber.New()

	database.ConnectDb()

	routes.SetUpRoutes(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
