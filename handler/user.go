package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devcamp-team-19/backend-sad/core/entity"
	"github.com/devcamp-team-19/backend-sad/core/module"
)

type UserHandler struct {
	productUc module.UserUsecase
}

func NewUserHandler(productUc module.UserUsecase) *UserHandler {
	return &UserHandler{
		productUc: productUc,
	}
}

func (hdl *UserHandler) GetAll(c *gin.Context) {
	Users, err := hdl.productUc.GetUsers(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Users})
}

func (hdl *UserHandler) GetSingle(c *gin.Context) {
	User, err := hdl.productUc.GetUser(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": User})
}

func (hdl *UserHandler) Create(c *gin.Context) {
	var input entity.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	// Create User
	User := entity.User{
		FullName: input.FullName,
		NIK:      input.NIK,
		Email:    input.Email,
		Address:  input.Address,
		Password: input.Password,
	}

	err := hdl.productUc.CreateUser(c, User)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "product succesfully created"})
}

func (hdl *UserHandler) Update(c *gin.Context) {
	var input entity.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	// Update User
	User := entity.User{
		FullName: input.FullName,
		NIK:      input.NIK,
		Email:    input.Email,
		Address:  input.Address,
		Password: input.Password,
	}

	err := hdl.productUc.UpdateUser(c, User)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "product succesfully updated"})
}
func (hdl *UserHandler) Delete(c *gin.Context) {
	err := hdl.productUc.DeleteUser(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "product succesfully deleted"})
}
