package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/helper"
	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/model/web"
	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/model/web/request"
	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/service"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (controller *OrderControllerImpl) Create(ctx *gin.Context) {
	var request request.OrderCreate
	err := ctx.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	response := controller.OrderService.Create(request)

	ctx.JSON(http.StatusOK, web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *OrderControllerImpl) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	request := request.OrderUpdate{}
	err := ctx.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	err = controller.OrderService.Update(request, id)
	helper.PanicIfError(err)

	ctx.JSON(http.StatusOK, web.Response{
		Code:   http.StatusOK,
		Status: "OK",
	})
}

func (controller *OrderControllerImpl) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := controller.OrderService.Delete(id)
	helper.PanicIfError(err)

	ctx.JSON(http.StatusOK, web.Response{
		Code:   http.StatusOK,
		Status: "OK",
	})
}

func (controller *OrderControllerImpl) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	order, err := controller.OrderService.FindById(id)
	helper.PanicIfError(err)

	ctx.JSON(http.StatusOK, web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   order,
	})
}

func (controller *OrderControllerImpl) FindAll(ctx *gin.Context) {
	orders, err := controller.OrderService.FindAll()
	helper.PanicIfError(err)

	ctx.JSON(http.StatusOK, web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   orders,
	})
}
