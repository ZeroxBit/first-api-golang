package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username        string    `json: "username" grom:"not null unique)" `
	FullName        string    `json: "fullname" gorm: "not null"`
	Email           string    `json: "email" grom:"unique" `
	Password        string    `json: "password,onmitempty" grom:"not null; type:varchar(255)" `
	ConfirmPassword string    `json: "confirmPassword,omitempty" gorm:"-"` // con el - no manda este valor a la vase de datos, lo omite
	Picture         string    `json:"picture"`
	Comment         []Comment `json:"comments,omitempty`
}
