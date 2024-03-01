package main

import (
	"latihan_api/app"
	"latihan_api/controller"
	"latihan_api/helper"
	"latihan_api/middleware"
	"latihan_api/repository"
	"latihan_api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepositoryImpl()
	userService := service.NewUserServiceImpl(userRepository, db, validate)
	userController := controller.NewUserControllerImpl(userService)
	router := app.NewRouter(userController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicError(err)
}
