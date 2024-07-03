package router

import (
    "github.com/gin-gonic/gin"
    "tugas-akhir/controllers"
    "tugas-akhir/middleware"
)

func CharacterRoutes(router *gin.Engine) {
    characterGroup := router.Group("/characters")

    characterGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
    {
        characterGroup.POST("/", controllers.CreateCharacter)
        characterGroup.GET("/", controllers.GetCharacters)
        characterGroup.GET("/:id", controllers.GetCharacter)
        characterGroup.PUT("/:id", controllers.UpdateCharacter)
        characterGroup.DELETE("/:id", controllers.DeleteCharacter)
    }
}