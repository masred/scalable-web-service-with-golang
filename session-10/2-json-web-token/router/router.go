package router

import (
	"github.com/gin-gonic/gin"
	"github.com/masred/scalable-web-service-with-golang/session-10/2-json-web-token/controllers"
	"github.com/masred/scalable-web-service-with-golang/session-10/2-json-web-token/middlewares"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)

		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
	}

	return r
}
