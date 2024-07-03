package router

import (
    "github.com/gin-gonic/gin"
    "tugas-akhir/controllers"
    "tugas-akhir/middleware"
)

func TeamCharacterRoutes(router *gin.Engine) {
    teamCharacterGroup := router.Group("/team-characters")

    teamCharacterGroup.Use(middleware.AuthMiddleware())
    {
        teamCharacterGroup.POST("/", controllers.CreateTeamCharacter)
        teamCharacterGroup.GET("/", controllers.GetTeamCharacters)
        teamCharacterGroup.GET("/:id", controllers.GetTeamCharacter)
        teamCharacterGroup.PUT("/:id", controllers.UpdateTeamCharacter)
        teamCharacterGroup.DELETE("/:id", controllers.DeleteTeamCharacter)
    }
}