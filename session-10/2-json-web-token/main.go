package main

import (
	"github.com/masred/scalable-web-service-with-golang/session-10/2-json-web-token/database"
	"github.com/masred/scalable-web-service-with-golang/session-10/2-json-web-token/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
