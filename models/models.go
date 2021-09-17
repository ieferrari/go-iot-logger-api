package models

import (
	"fmt"
	"time"

	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Book model
type User struct {
	gorm.Model

	Username string   `json:"Username" gorm:"unique" gorm:"primary_key"`
	Email    string   `json:"Email"`
	Pswd     string   `json:"Pswd"`
	Device   []Device `gorm:"foreignKey:UserID"`
}

type Device struct {
	gorm.Model
	UserID      uint
	Name        string   `json:"Name"`
	Description string   `json:"Description"`
	Uname_Dname string   `gorm:"unique" gorm:"primary_key"` //internal code: username-deviceName
	Sensor      []Sensor `gorm:"foreignKey:Device"`
}

type Sensor struct {
	gorm.Model
	Name              string
	Description       string
	Device            string     //Uname_Dname_Sname
	Uname_Dname_Sname string     `gorm:"unique" gorm:"primary_key"`
	Registry          []Registry `gorm:"foreignKey:Sensor"`
}

type Registry struct {
	gorm.Model
	Sensor    string //Uname_Dname_Sname
	Timestamp time.Time
	Value     float64
}

func Foo() {
	fmt.Println("bar")
}
