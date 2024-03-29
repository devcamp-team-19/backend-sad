package main

import (
	"github.com/fajartd02/mygallery/config"
	"github.com/fajartd02/mygallery/core/entity"
	"github.com/fajartd02/mygallery/core/module"
	"github.com/fajartd02/mygallery/handler"
	commentrepository "github.com/fajartd02/mygallery/repository/comment"
	filerepository "github.com/fajartd02/mygallery/repository/file"
	reportrepository "github.com/fajartd02/mygallery/repository/report"
	userrepository "github.com/fajartd02/mygallery/repository/user"
	uservoterepository "github.com/fajartd02/mygallery/repository/user-vote"
	"github.com/fajartd02/mygallery/routes"
)

func main() {

	db := config.InitDB()
	// cfg := config.InitConfig()
	db.AutoMigrate(&entity.User{}, &entity.Report{}, &entity.UserVote{}, &entity.Comment{}, &entity.File{})

	userRepo := userrepository.New()
	commentRepo := commentrepository.New()
	fileRepo := filerepository.New()
	reportRepo := reportrepository.New()
	userVoteRepo := uservoterepository.New()

	userUc := module.NewUserUsecase(userRepo)
	commentUc := module.NewCommentUsecase(commentRepo, userRepo)
	userVoteUc := module.NewUserVoteUsecase(userVoteRepo)
	fileUc := module.NewFileUseCase(fileRepo, reportRepo)
	reportUc := module.NewReportUseCase(reportRepo, userRepo)

	userHdl := handler.NewUserHandler(userUc)
	commentHdl := handler.NewCommentHandler(commentUc)
	userVoteHdl := handler.NewUserVoteHandler(userVoteUc)
	fileHdl := handler.NewFileHandler(fileUc)
	reportHdl := handler.NewReportHandler(reportUc)

	r := routes.SetupRoutes(db, *userHdl, *commentHdl, *fileHdl, *userVoteHdl, *reportHdl)

	r.Run(":8080")
}
