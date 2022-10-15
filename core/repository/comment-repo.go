package repository

import (
	"errors"

	"github.com/devcamp-team-19/backend-sad/core/entity"
	"github.com/gin-gonic/gin"
)

var ErrRecordCommentNotFound = errors.New("record report not found")

type CommentRepository interface {
	FindAll(c *gin.Context) ([]entity.Comment, error)
	Create(c *gin.Context) error
}
