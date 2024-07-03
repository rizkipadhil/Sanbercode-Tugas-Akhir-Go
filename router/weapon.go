package router

import (
    "github.com/gin-gonic/gin"
    "tugas-akhir/controllers"
    "tugas-akhir/middleware"
)

func WeaponRoutes(router *gin.Engine) {
    weaponGroup := router.Group("/weapons")

    weaponGroup.Use(middleware.AuthMiddleware())
    {
        weaponGroup.POST("/", controllers.CreateWeapon)
        weaponGroup.GET("/", controllers.GetWeapons)
        weaponGroup.GET("/:id", controllers.GetWeapon)
        weaponGroup.PUT("/:id", controllers.UpdateWeapon)
        weaponGroup.DELETE("/:id", controllers.DeleteWeapon)
    }
}