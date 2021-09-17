package handlers

import(
  "fmt"
  "github.com/gofiber/fiber/v2"
  // "go-iot-logger-api/queries"
   // "go-iot-logger-api/models"
   // "golang.org/x/crypto/bcrypt"
  // "go-iot-logger-api/database"
)

type nestedStruct struct {
	Field1 string `json:"field_1"`
	Field2 string `json:"field_2"`
	Field3 string `json:"field_3"`
}


type simple_response struct {
	Content   string `json:"content"`
	Additioal string `json:"additional"`
	Nested  nestedStruct `json:"nested"`
}

func Home(c *fiber.Ctx)error{
  response :=simple_response{"hola","mundo",nestedStruct{"1.12323","4.343","6.4445"}}
	fmt.Println("Hello response__")
	return c.JSON(response)
}

func Hello(c *fiber.Ctx)error{
  type response struct{
    Data string `json:"data"`
  }
  my_response := response{"alive"}
  return c.JSON(my_response)
}
