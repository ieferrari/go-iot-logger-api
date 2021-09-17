package queries

import (
	"fmt"
	"go-iot-logger-api/database"
	"go-iot-logger-api/models"
	"testing"
)

func TestCreateDevice(t *testing.T) {
	database.ConnectDb()
	var my_device models.Device
	my_device.UserID = uint(1)
	my_device.Name = "ejemplo2223"
	my_device.Description = "ejemplo de descripción"
	my_device.Uname_Dname = ("pedroo" + "_" + my_device.Name)

	response := CreateDevice(&my_device)
	//os.Remove("books.db")
	if response != "OK" {
		t.Errorf("test no pasado %d", 2)
	}
}
func TestCreateDevice2(t *testing.T) {
	//database.ConnectDb()
	var my_device models.Device
	my_device.UserID = uint(1)
	my_device.Name = "ejemplo222"
	my_device.Description = "ejemplo de descripción"
	my_device.Uname_Dname = ("pedro" + "_" + my_device.Name)

	response := CreateDevice(&my_device)
	//os.Remove("books.db")
	if response == "OK" {
		t.Errorf("test no pasado %d", 2)
	}
}

func TestGetDevice(t *testing.T) { // valid device
	Dname := "pedro"
	Uname := "ejemplo222"
	response, err := GetDevice(Dname, Uname)
	fmt.Println(response.Name)
	if err != nil {
		t.Error("not able to get device")
	}
}
func TestGetDevice2(t *testing.T) { // device does not exist
	Dname := "pedro"
	Uname := "ejemplo22"
	response, err := GetDevice(Dname, Uname)
	if response != nil && err != nil {
		t.Error("response and error should be nil")
	}
}

func TestEditDeviceDescription(t *testing.T) {
	Dname := "pedro"
	Uname := "ejemplo222"
	newDescription := "nueva"
	response := EditDeviceDescription(Dname, Uname, newDescription)
	updated_device, err := GetDevice(Dname, Uname)
	if err != nil {
		t.Error("query returned error: " + err.Error())
	}
	if response != "OK" {
		t.Error("query did not return OK ")
	}
	if updated_device.Description != newDescription {
		t.Error("did not update DB as expected")
	}
}

func TestDeleteDevice(t *testing.T) {
	Dname := "pedro"
	Uname := "ejemplo222"
	response := DeleteDevice(Dname, Uname)
	deleted_device, err := GetDevice(Dname, Uname)
	if response != "OK" {
		t.Error("query did not return OK ")
	}
	if deleted_device != nil && err != nil {
		t.Error("Device was not deleted")
	}
}

//dummy test to remove testing DB
// func TestClean(t *testing.T) {
// 	os.Remove("books.db")
// }
