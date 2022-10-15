package reportrepository

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/devcamp-team-19/backend-sad/core/entity"
	repository_intf "github.com/devcamp-team-19/backend-sad/core/repository"
)

var ErrDBNotFound = errors.New("repo: db not found")

type repository struct {
}

func (r repository) FindAll(c *gin.Context) ([]entity.Report, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) FindSingle(c *gin.Context, filename string) (entity.Report, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Create(c *gin.Context, report *entity.Report) (entity.Report, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Delete(c *gin.Context, filename string) error {
	//TODO implement me
	panic("implement me")
}

func New() repository_intf.ReportRepository {
	return &repository{}
}
