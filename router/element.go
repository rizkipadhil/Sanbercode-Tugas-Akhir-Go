package router

import (
    "github.com/gin-gonic/gin"
    "tugas-akhir/controllers"
    "tugas-akhir/middleware"
)

func ElementRoutes(router *gin.Engine) {
    elementGroup := router.Group("/elements")

    elementGroup.Use(middleware.AuthMiddleware())
    {
        elementGroup.POST("/", controllers.CreateElement)
        elementGroup.GET("/", controllers.GetElements)
        elementGroup.GET("/:id", controllers.GetElement)
        elementGroup.PUT("/:id", controllers.UpdateElement)
        elementGroup.DELETE("/:id", controllers.DeleteElement)
    }
}