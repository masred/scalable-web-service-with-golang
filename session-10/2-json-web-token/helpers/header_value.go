package helpers

import "github.com/gin-gonic/gin"

func GetContenType(ctx *gin.Context) string {
	return ctx.Request.Header.Get("Content-Type")
}
