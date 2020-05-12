package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"proxy-server/internal/config"
	"proxy-server/internal/proto/auth"
)

var (
	logger = config.SystemLogger
)

func Login(ctx *gin.Context) {
	var reqInput LoginData
	if err := ctx.ShouldBindJSON(&reqInput); err != nil {
		logger.Error(err)

		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": "invalid input"})

		return
	}

	loginProtoRequest := &auth.LoginRequest{
		Email:    reqInput.Email,
		Password: reqInput.Password,
	}

	if res, err := config.AuthClient.Login(ctx, loginProtoRequest); err == nil {
		config.SystemLogger.Debugln("response from server: " + res.Message)

		ctx.JSON(http.StatusOK, gin.H{"success": true})

		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "authentication service is temporary unavailable"})
	return

}

func Register(ctx *gin.Context) {
	var reqInput RegisterData
	if err := ctx.ShouldBindJSON(&reqInput); err != nil {
		logger.Error(err)

		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": "invalid input"})

		return
	}

	registerRequest := &auth.RegisterRequest{
		Email:     reqInput.Email,
		FirstName: reqInput.FirstName,
		LastName:  reqInput.LastName,
		Password:  reqInput.Password,
	}

	if res, err := config.AuthClient.Register(ctx, registerRequest); err == nil {
		config.SystemLogger.Debugln("response from server: " + res.Message)

		ctx.JSON(http.StatusOK, gin.H{"success": true})

		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "authentication service is temporary unavailable"})
	return
}
