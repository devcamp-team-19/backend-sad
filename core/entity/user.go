package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"fullName"`
	NIK      string `json:"nik"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type UserInput struct {
	FullName string `json:"fullName"`
	NIK      string `json:"nik"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}
