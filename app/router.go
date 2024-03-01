package app

import (
	"latihan_api/controller"
	"latihan_api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users/login", userController.Login)
	router.GET("/api/users/all", (userController.FindAll))
	router.POST("/api/users/register", userController.Register)
	router.GET("/api/users/logout", (userController.Logout))

	router.PanicHandler = exception.ErrorHandler

	return router
}
