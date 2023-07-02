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
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userVlidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
