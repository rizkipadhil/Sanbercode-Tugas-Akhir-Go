package router

import (
    "github.com/gin-gonic/gin"
    "tugas-akhir/controllers"
    "tugas-akhir/middleware"
)

func ActionRoutes(router *gin.Engine) {
    actionGroup := router.Group("/action")
    actionGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
    {
        actionGroup.PUT("/verify-team/:id", controllers.VerifyTeam)
    }
}