package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/masred/scalable-web-service-with-golang/session-08/3-practice/config"
	"github.com/masred/scalable-web-service-with-golang/session-08/3-practice/controllers"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/persons", inDB.GetPersons)
	router.POST("/persons", inDB.CreatePerson)
	router.GET("/persons/:id", inDB.GetPerson)
	router.PUT("/persons/:id", inDB.UpdatePerson)
	router.DELETE("/persons/:id", inDB.DeletePerson)
	router.Run(":3000")
}
