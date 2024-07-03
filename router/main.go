package router

import (
    "github.com/gin-gonic/gin"
    "tugas-akhir/controllers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Public routes
    router.GET("/", controllers.Greeting)
    AuthRoutes(router)

    ElementRoutes(router)
    WeaponRoutes(router)
    ArtifactRoutes(router)
    CharacterRoutes(router)

    return router
}