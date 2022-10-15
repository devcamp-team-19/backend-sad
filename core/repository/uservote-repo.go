package repository

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrRecordUserVoteNotFound = errors.New("record vote not found")

type UserVoteRepository interface {
	ChooseVotes(c *gin.Context) error
}
