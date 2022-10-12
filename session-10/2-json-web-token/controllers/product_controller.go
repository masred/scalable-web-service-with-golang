package controllers

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/masred/scalable-web-service-with-golang/session-10/2-json-web-token/database"
	"github.com/masred/scalable-web-service-with-golang/session-10/2-json-web-token/helpers"
	"github.com/masred/scalable-web-service-with-golang/session-10/2-json-web-token/models"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContenType(ctx)

	product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&product)
	} else {
		ctx.ShouldBind(&product)
	}

	product.UserID = userID

	err := db.Debug().Create(&product).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContenType(ctx)
	product := models.Product{}

	productId, _ := strconv.Atoi(ctx.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&product)
	} else {
		ctx.ShouldBind(&product)
	}

	product.UserID = userID
	product.ID = uint(productId)

	err := db.Model(&product).Where("id = ?", productId).Updates(models.Product{Title: product.Title, Description: product.Description}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
