package handlers

import (
	// "go-iot-logger-api/database"
	// "go-iot-logger-api/models"
	"time"
	"fmt"
  "go-iot-logger-api/queries"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	// "github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
  "os"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


// Login get user and password
func Login(c *fiber.Ctx) error {
	fmt.Println("LOGIN")
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var input LoginInput
	var ud UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	username := input.Username
	pass := input.Password

	user, err := queries.GetUserByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
	}

	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	}

  ud = UserData{
    ID:       user.ID,
    Username: user.Username,
    Email:    user.Email,
    Password: user.Pswd,
  }

	if !CheckPasswordHash(pass, ud.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Hour *72 ).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
