package main

import (
	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/app"
	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/controller"
	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/repository"
	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/service"
)

func main() {
	db := app.NewDB()
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(orderService)

	router := app.OrderRoute(orderController)

	router.Run()
}
