package http

import (
	"fmt"
	"tanneryApp/cmd/docs"
	"tanneryApp/internal/delivery/http/middleware"
	"tanneryApp/internal/infrastructure/config"
	"tanneryApp/internal/infrastructure/jwt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type AuthHandler interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

type UserHandler interface {
	GetCurrentProfile(ctx *gin.Context)
	CreateProfile(ctx *gin.Context)
}

func RegisterRoutes(router *gin.Engine, cfg *config.Config, jwtManager jwt.Manager, authHandler AuthHandler, userHandler UserHandler) {
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Http.Host, cfg.Http.Port)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", authHandler.SignUp)
		auth.POST("/sign-in", authHandler.SignIn)
	}

	user := api.Group("/user").Use(middleware.AuthMiddleware(jwtManager))
	{
		user.GET("/profile", userHandler.GetCurrentProfile)
		user.POST("/profile", userHandler.CreateProfile)
	}
}
