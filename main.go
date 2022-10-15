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
	"github.com/devcamp-team-19/backend-sad/routes"
)

func main() {

	db := config.Init()
	db.AutoMigrate(&entity.User{}, &entity.Report{}, &entity.UserVote{}, &entity.Comment{}, &entity.File{})

	userRepo := userrepository.New()
	commentRepo := commentrepository.New()
	fileRepo := filerepository.New()
	reportRepo := reportrepository.New()

	userUc := module.NewUserUsecase(userRepo)
	commentUc := module.NewCommentUsecase(commentRepo)
	fileUc := module.NewFileUseCase(fileRepo, reportRepo)

	userHdl := handler.NewUserHandler(userUc)
	commentHdl := handler.NewCommentHandler(commentUc)
	fileHdl := handler.NewFileHandler(fileUc)

	r := routes.SetupRoutes(db, *userHdl, *commentHdl, *fileHdl)

	r.Run(":8080")
}
