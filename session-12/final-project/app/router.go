package app

import (
	"github.com/gin-gonic/gin"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/controller"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/middleware"
)

func NewRoute(controller controller.UserController) *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.Register)
		userRouter.POST("/login", controller.Login)

		{
			userRouter.PUT("/", middleware.Auth(), controller.Update)
			userRouter.DELETE("/", middleware.Auth(), controller.Delete)
		}
	}

	return r
}
