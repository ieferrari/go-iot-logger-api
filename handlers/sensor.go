package handlers

import(
  "github.com/gofiber/fiber/v2"
  "go-iot-logger-api/queries"
)


func CreateSensor(c *fiber.Ctx)error{
  response := queries.CreateSensor()
  return c.SendString(response)
}

func EditSensor(c *fiber.Ctx)error{
  response := queries.EditSensor()
  return c.SendString(response)
}

func GetSensor(c *fiber.Ctx)error{
  response := queries.GetSensor()
  return c.SendString(response)
}

func LogSensor(c *fiber.Ctx)error{
  response := queries.LogSensor()
  return c.SendString(response)
}
