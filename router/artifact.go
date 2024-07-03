package router

import (
    "github.com/gin-gonic/gin"
    "tugas-akhir/controllers"
    "tugas-akhir/middleware"
)

func ArtifactRoutes(router *gin.Engine) {
    artifactGroup := router.Group("/artifacts")

    artifactGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
    {
        artifactGroup.POST("/", controllers.CreateArtifact)
        artifactGroup.GET("/", controllers.GetArtifacts)
        artifactGroup.GET("/:id", controllers.GetArtifact)
        artifactGroup.PUT("/:id", controllers.UpdateArtifact)
        artifactGroup.DELETE("/:id", controllers.DeleteArtifact)
    }
}