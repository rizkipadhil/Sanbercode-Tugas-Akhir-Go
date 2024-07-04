package router

import (
    "github.com/gin-gonic/gin"
    "tugas-akhir/controllers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Public routes
    router.GET("/", controllers.Greeting)
    // GetTeamsPublic
    router.GET("/teams-public", controllers.GetTeamsPublic)

    AuthRoutes(router)

    ElementRoutes(router)
    WeaponRoutes(router)
    ArtifactRoutes(router)
    CharacterRoutes(router)
    TeamRoutes(router)
    TeamCharacterRoutes(router)
    ActionRoutes(router)

    return router
}