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

    return router
}