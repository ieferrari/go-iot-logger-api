package queries

import(
  "fmt"
  //"github.com/gofiber/fiber/v2"
  "go-iot-logger-api/database"
  "go-iot-logger-api/models"
  "github.com/jinzhu/gorm"
    // "go-iot-logger-api/models"
    // "database/sql"
)

func CreateUser(new_user *models.User)string{
  db_connection := database.DBConn
	response :=db_connection.Create(&new_user)
  if response.Error != nil{
    return response.Error.Error()  // from response, get error, convert to string
  }
  return "OK"
}

func GetUserHash(id string)(string){
  db_connection := database.DBConn
  var user models.User
  db_connection.First(&user, id)
  return user.Pswd
}

func GetUserByUsername(username string) (*models.User, error) {
  fmt.Println("GetUserByUsername")
	db_connection := database.DBConn
	var user models.User
	err := db_connection.Where(&models.User{Username: username}).Find(&user).Error
  if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}



func DeleteUserByUsername(username string) (string, error){
  db_connection := database.DBConn
  var user models.User
  username = "Juan"
	err := db_connection.Unscoped().Where(&models.User{Username: username}).Delete(&user).Error
  //db_connection.Unscoped().Where(&models.User{Username: username}).Delete(&user)
  if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return (username+" not found, nothing deleted"), nil
		}
		return (username+" was not deleted, something happend"), err
	}
	return (username+" was successfully deleted"), nil
}
