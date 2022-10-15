package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/devcamp-team-19/backend-sad/handler"
)

func SetupRoutes(db *gorm.DB, userHdl handler.UserHandler) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/users", userHdl.GetAll)
	r.GET("/users/:id", userHdl.GetSingle)
	r.PUT("/users/:id", userHdl.Update)
	r.POST("/users", userHdl.Create)
	r.DELETE("/users/:id", userHdl.Delete)
	return r
}
