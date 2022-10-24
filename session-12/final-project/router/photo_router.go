package router

import (
	"github.com/gin-gonic/gin"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/app"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/controller"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/middleware"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/repository"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/service"
)

func PhotoRouter(router *gin.Engine) {
	db := app.NewDB()

	repo := repository.NewPhotoRepository(db)
	srv := service.NewPhotoService(repo)
	ctrl := controller.NewPhotoController(srv)

	photoRouter := router.Group("/photos", middleware.AuthMiddleware())

	{
		photoRouter.POST("/", ctrl.Create)
		photoRouter.GET("/", ctrl.GetAll)
		photoRouter.PUT("/:id", middleware.PhotoMiddleware(ctrl.PhotoService), ctrl.Update)
		photoRouter.DELETE("/:id", middleware.PhotoMiddleware(ctrl.PhotoService), ctrl.Delete)
	}
}
