package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/devcamp-team-19/backend-sad/handler"
)

func SetupRoutes(
	db *gorm.DB,
	userHdl handler.UserHandler,
	commentHdl handler.CommentHandler,
	userVoteHdl handler.UserVoteHandler) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/users", userHdl.GetAll)
		apiV1.GET("/users/:id", userHdl.GetSingle)
		apiV1.PUT("/users/:id", userHdl.Update)
		apiV1.POST("/users", userHdl.Create)
		apiV1.DELETE("/users/:id", userHdl.Delete)

		apiV1.POST("/reports/:reportId", commentHdl.Create)
		apiV1.GET("/reports/:reportId", commentHdl.GetAll)

		apiV1.POST("/reports/1/votes", userVoteHdl.VotingReport)
	}

	return r
}
