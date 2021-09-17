package handlers

import (
	"fmt"
	"go-iot-logger-api/handlers/utils"
	"go-iot-logger-api/models"
	"go-iot-logger-api/queries"

	"github.com/gofiber/fiber/v2"
)

func CreateDevice(c *fiber.Ctx) error {
	claims, _ := utils.ExtractTokenMetadata(c)
	username := claims.Username
	user_id := claims.UserID
	// check if the user requested is equel to JWT user
	if username != c.Params("user_name") {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Forbidden, requested user is different from JWT user",
		})
	}

	my_device := new(models.Device)
	if err := c.BodyParser(my_device); err != nil {
		fmt.Println("wrong parse")
		return c.Status(503).SendString(err.Error())
	}

	my_device.UserID = user_id
	my_device.Uname_Dname = (username + "_" + my_device.Name)

	response := queries.CreateDevice(my_device)
	return c.SendString(response)
}

func GetDevice(c *fiber.Ctx) error {
	claims, _ := utils.ExtractTokenMetadata(c)
	username := claims.Username
	//user_id := claims.UserID
	Dname := c.Params("dev_id")
	Uname := c.Params("user_name")
	// check if the user requested is equel to JWT user
	if username != Uname {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Forbidden, requested user is different from JWT user",
		})
	}

	response, err := queries.GetDevice(Dname, Uname)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	return c.JSON(response)
}

func EditDevice(c *fiber.Ctx) error {
	claims, _ := utils.ExtractTokenMetadata(c)
	username := claims.Username
	//user_id := claims.UserID
	Dname := c.Params("dev_id")
	Uname := c.Params("user_name")
	// check if the user requested is equel to JWT user
	if username != Uname {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Forbidden, requested user is different from JWT user",
		})
	}
	type NewDescription struct {
		Description string `json:"newDescription" xml:"newDescription" form:"newDescription"`
	}
	my_newDescription := new(NewDescription)
	if err := c.BodyParser(my_newDescription); err != nil {
		fmt.Println("could not parse newDescription for EditDevice")
		return c.Status(503).SendString(err.Error())
	}
	newDescription := my_newDescription.Description
	fmt.Println("newDescription: " + newDescription)

	response := queries.EditDeviceDescription(Dname, Uname, newDescription)
	return c.SendString(response)
}

func DeleteDevice(c *fiber.Ctx) error {
	claims, _ := utils.ExtractTokenMetadata(c)
	username := claims.Username
	//user_id := claims.UserID
	Dname := c.Params("dev_id")
	Uname := c.Params("user_name")
	// check if the user requested is equel to JWT user
	if username != Uname {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Forbidden, requested user is different from JWT user",
		})
	}
	response := queries.DeleteDevice(Dname, Uname)
	return c.SendString(response)
}

func GetDevicesForUser(c *fiber.Ctx) error {

	devices, err := queries.GetDevicesForUser("pedro")
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	return c.JSON(devices)
}
