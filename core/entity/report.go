package entity

import (
	"github.com/jinzhu/gorm"
)

type Report struct {
	gorm.Model
	UserVotes   []UserVote `json:"userVotes"`
	Comments    []Comment  `json:"comments"`
	UserID      uint
	EventTitle  string  `json:"eventTitle"`
	Category    uint    `json:"category"`
	ImageUrl    string  `json:"imageUrl"`
	Description string  `json:"description"`
	Longitude   string  `json"longitude"`
	Latitude    string  `json:"latitude"`
	radius      float64 `json:"radius"`
}

type ReportInput struct {
	UserVotes   []UserVoteInput `json:"userVotes"`
	Comments    []CommentInput  `json:"comments"`
	EventTitle  string          `json:"eventTitle"`
	Category    uint            `json:"category"`
	ImageUrl    string          `json:"imageUrl"`
	Description string          `json:"description"`
	Longitude   string          `json"longitude"`
	Latitude    string          `json:"latitude"`
	radius      float64         `json:"radius"`
}
