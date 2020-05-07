package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"msg": "not implemente"})
}

func Register(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"msg": "not implemente"})
}
