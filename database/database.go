package database

import (
	"fmt"
	"go-iot-logger-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var ( // env variables to load
	DBConn  *gorm.DB
	MESSAGE string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_NAME string
)

func ConnectDb() {
	// response := "host="//+DB_HOST+" user="+DB_USER+" password='"+DB_PASS+"' dbname="+DB_NAME+" port="+DB_PORT+" sslmode=disable"
	var err error
	DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
		panic(err)
	}
	fmt.Println("Connection Opened to Database")
	DBConn.AutoMigrate(
		&models.User{},
		&models.Device{},
		&models.Sensor{},
		&models.Registry{},
	)
}
