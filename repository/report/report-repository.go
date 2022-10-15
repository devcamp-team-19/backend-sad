package reportrepository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/devcamp-team-19/backend-sad/core/entity"
	repository_intf "github.com/devcamp-team-19/backend-sad/core/repository"
)

var ErrDBNotFound = errors.New("repo: db not found")

type repository struct {
}

func (r repository) FindAll(c *gin.Context) ([]entity.Report, error) {
	var reports []entity.Report

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to parse db to gorm")
	}

	if err := db.Model(&[]entity.Report{}).Preload("Comments").Preload("UserVotes").First(&reports).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrRecordReportNotFound
		}
		return nil, err
	}

	return reports, nil
}

func (r repository) FindSingle(c *gin.Context, reportID uint) (entity.Report, error) {
	report := entity.Report{}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return entity.Report{}, errors.New("failed to parse db to gorm")
	}

	if err := db.Model(&entity.Report{}).Preload("Comments").Preload("UserVotes").Where("id = ?", reportID).First(&report).Error; err != nil {
		return entity.Report{}, repository_intf.ErrRecordUserNotFound
	}

	return report, nil
}

func (r repository) Create(c *gin.Context, report entity.Report) (entity.Report, error) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return entity.Report{}, errors.New("failed to parse db to gorm")
	}

	if err := db.Create(&report).Error; err != nil {
		return entity.Report{}, errors.New("failed to create report")
	}

	return report, nil
}

func (r repository) Delete(c *gin.Context, filename string) error {
	//TODO implement me
	panic("implement me")
}

func New() repository_intf.ReportRepository {
	return &repository{}
}
