package repository

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/devcamp-team-19/backend-sad/core/entity"
)

var ErrRecordUserNotFound = errors.New("record not found")

type UserRepository interface {
	FindAll(c *gin.Context) ([]entity.User, error)
	FindSingle(c *gin.Context) (entity.User, error)
	Create(c *gin.Context, user entity.User) error
	Update(c *gin.Context, user entity.User) error
	Delete(c *gin.Context) error
}
