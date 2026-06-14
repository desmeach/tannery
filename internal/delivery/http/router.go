package http

import (
	"fmt"
	"tanneryApp/cmd/docs"
	"tanneryApp/internal/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type AuthHandler interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

func RegisterRoutes(router *gin.Engine, cfg *config.Config, authHandler AuthHandler) {
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Http.Host, cfg.Http.Port)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")

	users := api.Group("/auth")
	{
		users.POST("/sign-up", authHandler.SignUp)
		users.POST("/sign-in", authHandler.SignIn)
	}
}
