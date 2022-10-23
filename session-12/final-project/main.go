package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/app"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/controller"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/repository"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db := app.NewDB()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router := app.NewRoute(userController)
	router.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
}
