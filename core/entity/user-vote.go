package entity

import (
	"gorm.io/gorm"
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
