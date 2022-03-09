package server

import (
	middleware "be_entry_task/internal/http/handler"
	"be_entry_task/internal/http/handler/domain/auth/handler"
	user "be_entry_task/internal/http/handler/domain/user/handler"
	"github.com/julienschmidt/httprouter"
)

func ListRoute() *httprouter.Router {
	router := httprouter.New()

	//AUTH
	router.POST("/register", handler.NewRegister().Handle)
	router.POST("/login", handler.NewLogin().Handle)

	//FEATURE
	router.PUT("/users/profile", middleware.Auth(user.NewUpdateUser().Handle))
	router.PUT("/users/profile-picture", middleware.Auth(user.NewUpdatePicture().Handle))
	router.GET("/users/profile", middleware.Auth(user.NewGetProfile().Handle))

	return router
}
