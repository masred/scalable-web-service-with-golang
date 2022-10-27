package router

import (
	"github.com/gin-gonic/gin"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/app"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/controller"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/middleware"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/repository"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/service"
)

func SocialMediaRouter(router *gin.Engine) {
	db := app.NewDB()

	repo := repository.NewSocialMediaRepository(db)
	srv := service.NewSocialMediaService(repo)
	ctrl := controller.NewSocialMediaController(srv)

	socialMedia := router.Group("/socialmedias", middleware.AuthMiddleware())

	{
		socialMedia.GET("/", ctrl.GetAll)
		socialMedia.POST("/", ctrl.Create)
		{
			socialMedia.PUT("/:id", middleware.SocialMediaMiddleware(srv), ctrl.Update)
			socialMedia.DELETE("/:id", middleware.SocialMediaMiddleware(srv), ctrl.Delete)
		}
	}
}
