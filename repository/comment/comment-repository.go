package commentrepository

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/devcamp-team-19/backend-sad/core/entity"
	repository_intf "github.com/devcamp-team-19/backend-sad/core/repository"
)

type repositoryComment struct {
}

func New() repository_intf.CommentRepository {
	return &repositoryComment{}
}

// TODO: belom pake auth
func (r *repositoryComment) Create(c *gin.Context) error {
	var input entity.CommentInput
	var userId uint = 1 // TODO: dummy dulu, nanti isi ini pake jwt

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

	comment := entity.Comment{
		UserID:      userId,
		ReportID:    reportId,
		Description: input.Description,
	}

	if err := db.Create(&comment).Error; err != nil {
		return errors.New("failed to create comment")
	}

	return nil
}

// TODO: belom pake auth
func (r *repositoryComment) FindAll(c *gin.Context) ([]entity.Comment, error) {
	var comments []entity.Comment

	paramsId, err := strconv.ParseInt(c.Params.ByName("reportId"), 32, 32)
	if err != nil {
		return nil, errors.New("failed to convert params")
	}

	reportId := uint(paramsId)

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to parse db to gorm")
	}

	db.Raw("SELECT * FROM comments WHERE report_id = ?", reportId).Scan(&comments)
	if comments == nil {
		return nil, errors.New("comments not found")
	}
	return comments, nil
}
