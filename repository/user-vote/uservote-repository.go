package uservoterepository

import (
	repository_intf "github.com/devcamp-team-19/backend-sad/core/repository"
	"github.com/gin-gonic/gin"
)

type repositoryUserVote struct {
}

func New() repository_intf.UserVoteRepository {
	return &repositoryUserVote{}
}

func (r *repositoryUserVote) ChooseVotes(c *gin.Context) error {

	return nil
}
