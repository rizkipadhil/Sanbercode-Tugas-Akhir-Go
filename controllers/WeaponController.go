package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func validateWeapon(weapon *models.Weapon) (bool, string) {
    if weapon.Name == "" {
        return false, "Name is required"
    }
    return true, ""
}

func CreateWeapon(c *gin.Context) {
    var weapon models.Weapon
    if err := c.ShouldBindJSON(&weapon); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid data"})
        return
    }

    isValid, validationError := validateWeapon(&weapon)
    if !isValid {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": validationError})
        return
    }

    userName, _ := c.Get("username")
    weapon.CreatedBy = userName.(string)

    if result := database.DB.Create(&weapon); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to create weapon"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Weapon created successfully", "data": weapon})
}

func GetWeapons(c *gin.Context) {
    var weapons []models.Weapon
    if result := database.DB.Find(&weapons); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to retrieve weapons"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Weapons retrieved successfully", "data": weapons})
}

func GetWeapon(c *gin.Context) {
    var weapon models.Weapon
    id := c.Param("id")
    if result := database.DB.First(&weapon, id); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to retrieve weapon"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Weapon retrieved successfully", "data": weapon})
}

func UpdateWeapon(c *gin.Context) {
    var weapon models.Weapon
    id := c.Param("id")

    if err := c.ShouldBindJSON(&weapon); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid data"})
        return
    }

    isValid, validationError := validateWeapon(&weapon)
    if !isValid {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": validationError})
        return
    }

    userName, _ := c.Get("username")
    weapon.UpdatedBy = userName.(string)

    if result := database.DB.Model(&weapon).Where("id = ?", id).Updates(weapon); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to update weapon"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Weapon updated successfully", "data": weapon})
}

func DeleteWeapon(c *gin.Context) {
    id := c.Param("id")
    if result := database.DB.Delete(&models.Weapon{}, id); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to delete weapon"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Weapon deleted successfully"})
}