package module

import (
	"errors"
	"fmt"

	"github.com/devcamp-team-19/backend-sad/core/repository"
	"github.com/gin-gonic/gin"
)

type CommentUsecase interface {
	CreateComment(c *gin.Context) error
}

type commentUsecase struct {
	commentRepo repository.CommentRepository
}

var ErrCommentNotFound = errors.New("user comment: ")

func NewCommentUsecase(repo repository.CommentRepository) CommentUsecase {
	return &commentUsecase{repo}
}

func (em *commentUsecase) CreateComment(c *gin.Context) error {
	err := em.commentRepo.Create(c)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return nil
}
