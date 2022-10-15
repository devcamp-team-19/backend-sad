package userrepository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/devcamp-team-19/backend-sad/core/entity"
	repository_intf "github.com/devcamp-team-19/backend-sad/core/repository"
)

type repositoryUser struct {
}

func New() repository_intf.UserRepository {
	return &repositoryUser{}
}

func (r *repositoryUser) FindAll(c *gin.Context) ([]entity.User, error) {
	var users []entity.User

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to parse db to gorm")
	}

	err := db.Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrRecordUserNotFound
		}
		return nil, err
	}

	return users, nil
}

func (r *repositoryUser) FindSingle(c *gin.Context) (entity.User, error) {
	user := entity.User{}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return entity.User{}, errors.New("failed to parse db to gorm")
	}

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		return entity.User{}, repository_intf.ErrRecordUserNotFound
	}

	return user, nil
}

func (r *repositoryUser) Create(c *gin.Context) error {
	var input entity.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return errors.New("failed to create user")
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	// Create User
	User := entity.User{
		FullName: input.FullName,
		NIK:      input.NIK,
		Email:    input.Email,
		Address:  input.Address,
		Password: input.Password,
	}

	if err := db.Create(&User).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (r *repositoryUser) Update(c *gin.Context) error {
	var input entity.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return errors.New("failed to update user")
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	// Update User
	User := entity.User{
		FullName: input.FullName,
		NIK:      input.NIK,
		Email:    input.Email,
		Address:  input.Address,
		Password: input.Password,
	}

	if err := db.Where("id = ?", c.Param("id")).Save(&User).Error; err != nil {
		return errors.New("failed to update user")
	}

	return nil
}

func (r *repositoryUser) Delete(c *gin.Context) error {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	var user entity.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		return errors.New("failed to delete user")
	}

	if err := db.Delete(&user).Error; err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}
