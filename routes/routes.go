package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/devcamp-team-19/backend-sad/config"
	"github.com/devcamp-team-19/backend-sad/core/module"
	"github.com/devcamp-team-19/backend-sad/handler"
)

func SetupRoutes(
	db *gorm.DB,
	cfg config.Config,
	userHdl handler.UserHandler,
	commentHdl handler.CommentHandler,
	fileHdl handler.FileHandler,
	userVoteHdl handler.UserVoteHandler) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("host", cfg.Host)
	})

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "okss")
	})

	apiV1 := r.Group("/api/v1")
	{
		// Authentication and Authorization routes
		apiV1.POST("/login", userHdl.Login)
		apiV1.POST("/register", userHdl.Register)

		// User routes
		user := apiV1.Group("/users")
		user.Use(module.IsAuthorized())
		user.GET("", userHdl.GetAll)
		user.GET("/:id", userHdl.GetSingle)
		user.PUT("/:id", userHdl.Update)
		user.DELETE("/:id", userHdl.Delete)

		// File routes
		file := apiV1.Group("/files")
		file.POST("", fileHdl.UploadFile)
		file.StaticFS("/static", http.Dir("images"))

		// Comment routes
		comments := apiV1.Group("/comments")
		comments.Use(module.IsAuthorized())
		comments.POST("/:reportId", commentHdl.Create)
		comments.GET("/:reportId", commentHdl.GetAll)

		// UserVote routes
		reports := apiV1.Group("/reports")
		reports.POST("/:reportId/votes", userVoteHdl.VotingReport)
		reports.GET("/:reportId/votes", userVoteHdl.GetVotesInReport)
	}

	return r
}
