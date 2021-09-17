package queries

import (
	"fmt"
	"go-iot-logger-api/database"
	"go-iot-logger-api/models"

	"github.com/jinzhu/gorm"
)

func CreateDevice(new_device *models.Device) string {
	db_connection := database.DBConn
	response := db_connection.Create(&new_device)
	if response.Error != nil {
		fmt.Println("CreateDevice Failed!")
		return response.Error.Error()
	}
	return "OK"
}

func GetDevice(Dname string, Uname string) (*models.Device, error) {
	uname_dname := (Uname + "_" + Dname)
	fmt.Println(uname_dname)
	db_connection := database.DBConn
	var my_device models.Device
	err := db_connection.Where(&models.Device{Uname_Dname: uname_dname}).Find(&my_device).Error
	if err != nil {
		fmt.Println(err.Error())
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println(uname_dname)
	return &my_device, nil

}

func EditDeviceDescription(Dname string, Uname string, newDescription string) string {
	uname_dname := (Uname + "_" + Dname)
	db_connection := database.DBConn
	var my_device models.Device
	err := db_connection.Where(&models.Device{Uname_Dname: uname_dname}).Find(&my_device).Error
	if err != nil {
		fmt.Println("error, could not edit: " + uname_dname)
		return err.Error()
	}
	my_device.Description = newDescription
	db_connection.Save(&my_device)
	return "OK"
}

func DeleteDevice(Dname string, Uname string) string {
	uname_dname := (Uname + "_" + Dname)
	db := database.DBConn
	var my_device models.Device
	err := db.Where(&models.Device{Uname_Dname: uname_dname}).Find(&my_device).Error
	if err != nil {
		return err.Error()
	}
	db.Unscoped().Delete(&my_device)
	return "OK"
}

func GetDevicesForUser(username string) (models.User, error) {
	db := database.DBConn
	var user models.User
	var devices models.Device
	err := db.Where(&models.User{Username: username}).Find(&user).Error

	db.Model(&user).Association("Device").Find(&user.Device)
	fmt.Println(devices)
	return user, err
}
