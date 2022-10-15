package main

import (
	"github.com/devcamp-team-19/backend-sad/config"
	"github.com/devcamp-team-19/backend-sad/core/entity"
	"github.com/devcamp-team-19/backend-sad/core/module"
	"github.com/devcamp-team-19/backend-sad/handler"
	productrepository "github.com/devcamp-team-19/backend-sad/repository"
	"github.com/devcamp-team-19/backend-sad/routes"
)

func main() {

	db := config.Init()
	db.AutoMigrate(&entity.User{})

	userRepo := productrepository.New()
	userUc := module.NewUserUsecase(userRepo)
	userHdl := handler.NewUserHandler(userUc)

	r := routes.SetupRoutes(db, *userHdl)
	r.Run(":8080")
}
