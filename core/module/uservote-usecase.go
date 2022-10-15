package module

import (
	"errors"
	"fmt"

	"github.com/devcamp-team-19/backend-sad/core/repository"
	"github.com/gin-gonic/gin"
)

type UserVoteUsecase interface {
	VotingReport(c *gin.Context) error
}

type userVoteUsecase struct {
	userRepo repository.UserVoteRepository
}

var ErrUserVoteNotFound = errors.New("uservote error: ")

func NewUserVoteUsecase(repo repository.UserVoteRepository) UserVoteUsecase {
	return &userVoteUsecase{repo}
}

func (em *userVoteUsecase) VotingReport(c *gin.Context) error {
	err := em.userRepo.ChooseVotes(c)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCommentNotFound, err)
	}
	return nil
}
