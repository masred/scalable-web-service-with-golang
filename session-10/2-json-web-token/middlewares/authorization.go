package middlewares

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/masred/scalable-web-service-with-golang/session-10/2-json-web-token/database"
	"github.com/masred/scalable-web-service-with-golang/session-10/2-json-web-token/models"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(ctx.Param("productId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		product := models.Product{}

		err = db.Select("user_id").First(&product, uint(productId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if product.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorizad",
				"message": "you are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}
