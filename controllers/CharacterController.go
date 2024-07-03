package controllers

import (
    "net/http"
    "tugas-akhir/database"
    "tugas-akhir/models"
    "github.com/gin-gonic/gin"
)

func validateCharacter(character models.Character) []string {
    var errors []string
    if character.Name == "" {
        errors = append(errors, "Name is required")
    }
    if character.ElementID == 0 {
        errors = append(errors, "Element ID is required")
    }
    if character.WeaponID == 0 {
        errors = append(errors, "Weapon ID is required")
    }
    if character.Rarity < 1 || character.Rarity > 5 {
        errors = append(errors, "Rarity must be between 1 and 5")
    }
    return errors
}

func CreateCharacter(c *gin.Context) {
    var character models.Character
    if err := c.ShouldBindJSON(&character); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid data"})
        return
    }

    if validationErrors := validateCharacter(character); len(validationErrors) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": validationErrors})
        return
    }

    if result := database.DB.Create(&character); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Creation failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Character created successfully"})
}

func GetCharacter(c *gin.Context) {
    id := c.Param("id")
    var character models.Character
    if result := database.DB.Preload("Element").Preload("Weapon").First(&character, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Character not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "data": character})
}

func GetCharacters(c *gin.Context) {
    var characters []models.Character
    database.DB.Preload("Element").Preload("Weapon").Find(&characters)
    c.JSON(http.StatusOK, gin.H{"error": false, "data": characters})
}

func UpdateCharacter(c *gin.Context) {
    id := c.Param("id")
    var character models.Character
    if result := database.DB.First(&character, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Character not found"})
        return
    }

    if err := c.ShouldBindJSON(&character); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid data"})
        return
    }

    if validationErrors := validateCharacter(character); len(validationErrors) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": validationErrors})
        return
    }

    if result := database.DB.Save(&character); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Update failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Character updated successfully"})
}

func DeleteCharacter(c *gin.Context) {
    id := c.Param("id")
    var character models.Character
    if result := database.DB.First(&character, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Character not found"})
        return
    }

    if result := database.DB.Delete(&character, id); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Delete failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Character deleted successfully"})
}