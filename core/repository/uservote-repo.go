package repository

import (
	"errors"

	"github.com/devcamp-team-19/backend-sad/core/entity"
	"github.com/gin-gonic/gin"
)

var ErrRecordUserVoteNotFound = errors.New("record vote not found")

type UserVoteRepository interface {
	ChooseVotes(c *gin.Context) error
	GetVotesInReport(c *gin.Context) (entity.Votes, error)
}
