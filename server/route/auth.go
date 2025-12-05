package route

import (
	"server/handler"
	"server/middleware"
	"server/service"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authHandler := GenerateAuthHandler()

	authRoute := router.Group("/auth")
	{
		authRoute.POST("login", authHandler.Login)
		authRoute.POST("logout", authHandler.Logout)
	}

	userRoute := router.Group("/user")
	{
		userRoute.Use(middleware.JwtAuthMiddleware())
		userRoute.GET("", authHandler.GetUser)
	}
}

func GenerateAuthHandler() *handler.AuthHandler {
	service := service.NewAuthService()
	handler := handler.NewAuthHandler(service)

	return handler
}
