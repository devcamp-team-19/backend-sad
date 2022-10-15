package main

import (
	"github.com/devcamp-team-19/backend-sad/config"
	"github.com/devcamp-team-19/backend-sad/core/entity"
	"github.com/devcamp-team-19/backend-sad/core/module"
	"github.com/devcamp-team-19/backend-sad/handler"
	commentrepository "github.com/devcamp-team-19/backend-sad/repository/comment"
	filerepository "github.com/devcamp-team-19/backend-sad/repository/file"
	reportrepository "github.com/devcamp-team-19/backend-sad/repository/report"
	userrepository "github.com/devcamp-team-19/backend-sad/repository/user"
	uservoterepository "github.com/devcamp-team-19/backend-sad/repository/user-vote"
	"github.com/devcamp-team-19/backend-sad/routes"
)

func main() {

	db := config.InitDB()
	cfg := config.InitConfig()
	db.AutoMigrate(&entity.User{}, &entity.Report{}, &entity.UserVote{}, &entity.Comment{}, &entity.File{})

	userRepo := userrepository.New()
	commentRepo := commentrepository.New()
	fileRepo := filerepository.New()
	reportRepo := reportrepository.New()
	userVoteRepo := uservoterepository.New()

	userUc := module.NewUserUsecase(userRepo)
	commentUc := module.NewCommentUsecase(commentRepo)
	userVoteUc := module.NewUserVoteUsecase(userVoteRepo)
	fileUc := module.NewFileUseCase(fileRepo, reportRepo)

	userHdl := handler.NewUserHandler(userUc)
	commentHdl := handler.NewCommentHandler(commentUc)
	userVoteHdl := handler.NewUserVoteHandler(userVoteUc)
	fileHdl := handler.NewFileHandler(fileUc)

	r := routes.SetupRoutes(db, cfg, *userHdl, *commentHdl, *fileHdl, *userVoteHdl)

	r.Run(":8080")
}
