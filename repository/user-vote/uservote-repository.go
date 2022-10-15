package uservoterepository

import (
	"fmt"
	"log"
	"strconv"

	repository_intf "github.com/devcamp-team-19/backend-sad/core/repository"
	"github.com/gin-gonic/gin"
)

type repositoryUserVote struct {
}

func New() repository_intf.UserVoteRepository {
	return &repositoryUserVote{}
}

func (r *repositoryUserVote) ChooseVotes(c *gin.Context) error {
	query := c.Query("is_voting")
	isVoting, err := strconv.ParseBool(query)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %t\n", query, isVoting)

	return nil
}
