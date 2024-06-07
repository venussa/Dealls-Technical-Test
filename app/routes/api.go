package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/venussa/Dealls-Technical-Test/handlers"
)

func SetupAuthRoutes(router *gin.Engine) {
    auth := router.Group("/auth/")
    {
        auth.POST("/register", handlers.Register)
        auth.POST("/login", handlers.Login)
    }
}