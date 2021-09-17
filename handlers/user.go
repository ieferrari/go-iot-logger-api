package handlers

import (
	// "fmt"
	"go-iot-logger-api/models"
	"go-iot-logger-api/queries"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	// "fmt"
	"go-iot-logger-api/handlers/utils"
	"time"
	// "go-iot-logger-api/database"
)

type my_error struct {
	My_error string `json:"my_error"`
	Error    string `json:"error"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// is_valid := CheckPasswordHash(password, hash)

func CreateUser(context *fiber.Ctx) error {
	new_user := new(models.User)
	if err := context.BodyParser(new_user); err != nil {
		return context.Status(503).SendString(err.Error())
	}

	hash, err := hashPassword(new_user.Pswd)
	if err != nil {
		my_error := new(my_error)
		my_error.My_error = "failed to make a hash for password"
		my_error.Error = err.Error()
		return context.JSON(my_error)
	}
	new_user.Pswd = hash

	creation := queries.CreateUser(new_user)
	if creation != "OK" {
		return context.Status(503).SendString(creation)
	}
	return context.JSON(new_user)
}

// Input fiber.context with bearer: jwt and url param: username
// chek: if the url username is not equal to the Username in jwt, return error
// queries the database, calling queries.GetUserByUsername
// check: if no error in query, return error
// check: if user does not exist, return error
// return JSON with Username Email and CreatedAt
func GetUser(c *fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(c)
	username := claims.Username
	// check if the user requested is equel to JWT user
	if username != c.Params("user_name") {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Forbidden, requested user is different from JWT user",
		})
	}

	user, err := queries.GetUserByUsername(username)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if user == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   "Error, apparently we could not find you on the database",
		})
	}
	type user_simple struct {
		Username  string    `json:"Username"`
		Email     string    `json:"Email"`
		CreatedAt time.Time `json:"CreatedAt"`
	}

	var my_user_simple user_simple
	my_user_simple.Username = user.Username
	my_user_simple.Email = user.Email
	my_user_simple.CreatedAt = user.CreatedAt

	return c.Status(200).JSON(user) //my_user_simple)
}

func DeleteUser(c *fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(c)
	username := claims.Username
	// check if the user requested is equel to JWT user
	if username != c.Params("user_name") {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Forbidden, requested user is different from JWT user",
		})
	}

	response, err := queries.DeleteUserByUsername(username)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"msg": response,
	})
}
