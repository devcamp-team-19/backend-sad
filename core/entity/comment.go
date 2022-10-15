package entity

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserID      uint
	ReportID    uint
	Description string `json:"description"`
}

type CommentInput struct {
	Description string `json:"description"`
}
