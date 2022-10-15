package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/devcamp-team-19/backend-sad/core/module"
	"github.com/devcamp-team-19/backend-sad/handler"
)

func SetupRoutes(db *gorm.DB, userHdl handler.UserHandler, commentHdl handler.CommentHandler) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "okss")
	})

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/login", userHdl.Login)
		apiV1.POST("/register", userHdl.Register)

		user := apiV1
		user.Use(module.IsAuthorized())
		user.GET("/users", userHdl.GetAll)
		user.GET("/users/:id", userHdl.GetSingle)
		user.PUT("/users/:id", userHdl.Update)
		user.DELETE("/users/:id", userHdl.Delete)

		apiV1.POST("/reports/:reportId", commentHdl.Create)
		apiV1.GET("/reports/:reportId", commentHdl.GetAll)
	}

	return r
}
