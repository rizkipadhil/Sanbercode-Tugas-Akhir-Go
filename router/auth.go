package router

import (
    "github.com/gin-gonic/gin"
    "tugas-akhir/controllers/auth"
    "tugas-akhir/middleware"
)

func AuthRoutes(router *gin.Engine) {
    authGroup := router.Group("/auth")
    authGroup.POST("/login", auth.Login)
    authGroup.POST("/register", auth.Register)

    authGroup.Use(middleware.AuthMiddleware())
    {
        authGroup.POST("/refresh_token", auth.RefreshToken)
        authGroup.POST("/logout", auth.Logout)
        authGroup.GET("/user", auth.UserAuth)
        authGroup.POST("/change_password", auth.ChangePassword)
    }
}