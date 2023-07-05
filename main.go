package main

import (
	"bulletin-board-rest-api/controller"
	"bulletin-board-rest-api/db"
	"bulletin-board-rest-api/repository"
	"bulletin-board-rest-api/router"
	"bulletin-board-rest-api/usecase"
	"bulletin-board-rest-api/validator"
)

func main() {
	db := db.NewDB()
	userVlidator := validator.NewUserValidator()
	questValidator := validator.NewQuestValidator()
	userRepository := repository.NewUserRepository(db)
	questRepository := repository.NewQuestRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userVlidator)
	questUsecase := usecase.NewQuestUsecase(questRepository, userRepository, questValidator)
	userController := controller.NewUserController(userUsecase)
	questController := controller.NewQuestController(questUsecase)
	e := router.NewRouter(userController, questController)
	e.Logger.Fatal(e.Start(":8000"))
}
