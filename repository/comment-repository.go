package repository

import (
	"errors"
	"strconv"

	"github.com/devcamp-team-19/backend-sad/core/entity"
	repository_intf "github.com/devcamp-team-19/backend-sad/core/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type repositoryComment struct {
}

func NewReport() repository_intf.CommentRepository {
	return &repository{}
}

func (r *repositoryComment) Create(c *gin.Context) error {
	var input entity.CommentInput
	var userId uint = 2 // dummy dulu, nanti isi ini pake jwt auth

	paramsId, err := strconv.ParseInt(c.Params.ByName("reportId"), 32, 32)
	if err != nil {
		return errors.New("failed to convert params")
	}
	reportId := uint(paramsId)

	if err := c.ShouldBindJSON(&input); err != nil {
		return errors.New("failed to create comment")
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	Comment := entity.Comment{
		UserID:      userId,
		ReportID:    reportId,
		Description: input.Description,
	}

	if err := db.Create(&Comment).Error; err != nil {
		return errors.New("failed to create comment")
	}

	return nil
}
