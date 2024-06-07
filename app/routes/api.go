package routes

import (
    "github.com/gin-gonic/gin"
    "login-api/handlers"
)

func SetupAuthRoutes(router *gin.Engine) {
    auth := router.Group("/auth/")
    {
        auth.POST("/register", handlers.Register)
        auth.POST("/login", handlers.Login)
    }
}