package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Reports   []Report   `json:"reports"`
	UserVotes []UserVote `json:"userVotes"`
	Comments  []Comment  `json:"comments"`
	FullName  string     `json:"fullName"`
	NIK       string     `json:"nik"`
	Email     string     `json:"email"`
	Address   string     `json:"address"`
	Password  string     `json:"password"`
}

type UserInput struct {
	FullName string `json:"fullName"`
	NIK      string `json:"nik"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
