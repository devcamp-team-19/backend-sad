package repository

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrRecordCommentNotFound = errors.New("record report not found")

type CommentRepository interface {
	Create(c *gin.Context) error
}
