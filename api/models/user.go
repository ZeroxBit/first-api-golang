package models

import (
	"github.com/jinzhu/gorm"
)

// User usuario del sistema
type User struct {
	gorm.Model
	Username        string    `json: "username" gorm:"not null; unique" `
	FullName        string    `json: "fullname" gorm: "not null"`
	Email           string    `json: "email" gorm:"not null; unique"`
	Password        string    `json: "password,onmitempty" gorm:"not null; type:varchar(255)" `
	ConfirmPassword string    `json: "confirmPassword,omitempty" gorm:"-"` // con el - no manda este valor a la vase de datos, lo omite
	Picture         string    `json:"picture"`
	Comment         []Comment `json:"comments,omitempty`
}
