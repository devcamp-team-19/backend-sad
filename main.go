package main

import (
	"github.com/devcamp-team-19/backend-sad/config"
	"github.com/devcamp-team-19/backend-sad/core/entity"
	"github.com/devcamp-team-19/backend-sad/core/module"
	"github.com/devcamp-team-19/backend-sad/handler"
	commentrepository "github.com/devcamp-team-19/backend-sad/repository/comment"
	userrepository "github.com/devcamp-team-19/backend-sad/repository/user"
	uservoterepository "github.com/devcamp-team-19/backend-sad/repository/user-vote"
	"github.com/devcamp-team-19/backend-sad/routes"
)

func main() {

	db := config.Init()
	db.AutoMigrate(&entity.User{}, &entity.Report{}, &entity.UserVote{}, &entity.Comment{})

	userRepo := userrepository.New()
	userUc := module.NewUserUsecase(userRepo)
	userHdl := handler.NewUserHandler(userUc)

	commentRepo := commentrepository.New()
	commentUc := module.NewCommentUsecase(commentRepo)
	commentHdl := handler.NewCommentHandler(commentUc)

	userVoteRepo := uservoterepository.New()
	userVoteUc := module.NewUserVoteUsecase(userVoteRepo)
	userVoteHdl := handler.NewUserVoteHandler(userVoteUc)

	r := routes.SetupRoutes(db, *userHdl, *commentHdl, *userVoteHdl)

	r.Run(":8080")
}
