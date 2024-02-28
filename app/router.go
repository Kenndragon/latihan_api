package app

import (
	"latihan_api/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users/login", userController.Login)
	router.POST("/api/users/register", userController.Register)

	// router.PanicHandler = exception.ErrorHandler

	return router
}
