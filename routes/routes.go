package routes

import (
	"go-iot-logger-api/handlers"
	"go-iot-logger-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/hello", handlers.Hello) //test if API is working

	app.Post("/new_user", handlers.CreateUser)
	app.Post("/login", handlers.Login)
	app.Get("/u/:user_name", middleware.Protected(), handlers.GetUser)
	app.Delete("/u/:user_name", middleware.Protected(), handlers.DeleteUser)

	app.Post("/u/:user_name/device", middleware.Protected(), handlers.CreateDevice)
	app.Get("/u/:user_name/device", middleware.Protected(), handlers.GetDevicesForUser)
	app.Get("/u/:user_name/device/:dev_id", middleware.Protected(), handlers.GetDevice)
	app.Put("/u/:user_name/device/:dev_id", middleware.Protected(), handlers.EditDevice)
	app.Delete("/u/:user_name/device/:dev_id", middleware.Protected(), handlers.DeleteDevice)

	app.Post("/u/:user_name/device/:dev_id/sensor", middleware.Protected(), handlers.CreateSensor)
	app.Put("/u/:user_name/device/:dev_id/sensor/:sensor_id", middleware.Protected(), handlers.EditSensor)
	app.Get("/u/:user_name/device/:dev_id/sensor/:sensor_id", middleware.Protected(), handlers.GetSensor)
	app.Post("/u/:user_name/device/:dev_id/sensor/:sensor_id", middleware.Protected(), handlers.LogSensor)
}
