package module

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/devcamp-team-19/backend-sad/core/entity"
	"github.com/devcamp-team-19/backend-sad/core/repository"
)

type UserUsecase interface {
	GetUsers(c *gin.Context) ([]entity.User, error)
	GetUser(c *gin.Context) (entity.User, error)
	CreateUser(c *gin.Context) error
	UpdateUser(c *gin.Context) error
	DeleteUser(c *gin.Context) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase use for initiate new user usecase
func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

var ErrUserNotFound = errors.New("user error: ")

func (em *userUsecase) GetUsers(c *gin.Context) ([]entity.User, error) {
	data, err := em.userRepo.FindAll(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return nil, fmt.Errorf("%w.", ErrUserNotFound)
		}
		return nil, fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return data, nil
}

func (em *userUsecase) GetUser(c *gin.Context) (entity.User, error) {
	data, err := em.userRepo.FindSingle(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return entity.User{}, fmt.Errorf("%w.", ErrUserNotFound)
		}
		return entity.User{}, fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return data, nil
}

func (em *userUsecase) CreateUser(c *gin.Context) error {
	err := em.userRepo.Create(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return fmt.Errorf("%w.", ErrUserNotFound)
		}
		return fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return nil
}

func (em *userUsecase) UpdateUser(c *gin.Context) error {
	err := em.userRepo.Update(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return fmt.Errorf("%w.", ErrUserNotFound)
		}
		return fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return nil
}

func (em *userUsecase) DeleteUser(c *gin.Context) error {
	err := em.userRepo.Delete(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return fmt.Errorf("%w.", ErrUserNotFound)
		}
		return fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return nil
}
