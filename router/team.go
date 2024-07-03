package router

import (
    "github.com/gin-gonic/gin"
    "tugas-akhir/controllers"
    "tugas-akhir/middleware"
)

func TeamRoutes(router *gin.Engine) {
    teamGroup := router.Group("/teams")

    teamGroup.Use(middleware.AuthMiddleware())
    {
        teamGroup.POST("/", controllers.CreateTeam)
        teamGroup.GET("/", controllers.GetTeams)
        teamGroup.GET("/:id", controllers.GetTeam)
        teamGroup.PUT("/:id", controllers.UpdateTeam)
        teamGroup.DELETE("/:id", controllers.DeleteTeam)
    }
}