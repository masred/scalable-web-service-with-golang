package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masred/scalable-web-service-with-golang/session-10/2-json-web-token/helpers"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}
		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
