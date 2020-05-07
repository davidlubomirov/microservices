package main

import (
	"github.com/gin-gonic/gin"
	"os"

	"proxy-server/internal/config"
	"proxy-server/internal/http/routes"
)

func main() {
	logger := config.SetupLogger()

	if _, err := config.SetupServerConfiguration(); err != nil {
		logger.Error(err)
		os.Exit(config.ExitInvalidEnvironmentConfiguration)
	}

	logger.Infoln("setting proxy-server routes")
	router := gin.Default()

	proxy_v1 := router.Group("/api/v1")
	{
		proxy_v1.GET("/status", routes.Status)
		proxy_v1.POST("/auth/login", routes.Login)
		proxy_v1.POST("/auth/register", routes.Login)
	}

	router.Run(":8080")
}
