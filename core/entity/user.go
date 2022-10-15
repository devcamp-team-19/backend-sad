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
	Reports   []ReportInput   `json:"reports"`
	UserVotes []UserVoteInput `json:"userVotes"`
	Comments  []CommentInput  `json:"comments"`
	FullName  string          `json:"fullName"`
	NIK       string          `json:"nik"`
	Email     string          `json:"email"`
	Address   string          `json:"address"`
	Password  string          `json:"password"`
}
