package controller

import (
	"github.com/gin-gonic/gin"
)

type CommentController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
