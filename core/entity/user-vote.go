package entity

import (
	"github.com/jinzhu/gorm"
)

type UserVote struct {
	gorm.Model
	UserID   uint
	ReportID uint
	IsUpVote bool `json:"isUpVote"`
}

type UserVoteInput struct {
	IsUpVote bool `json:"isUpVote"`
}
