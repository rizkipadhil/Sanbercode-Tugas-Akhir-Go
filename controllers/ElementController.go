package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func validateElement(element *models.Element) (bool, string) {
    if element.Name == "" {
        return false, "Name is required"
    }
    return true, ""
}

func CreateElement(c *gin.Context) {
    var element models.Element
    if err := c.ShouldBindJSON(&element); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid data"})
        return
    }

    isValid, validationError := validateElement(&element)
    if !isValid {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": validationError})
        return
    }

    if result := database.DB.Create(&element); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to create element"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Element created successfully", "data": element})
}

func GetElements(c *gin.Context) {
    var elements []models.Element
    if result := database.DB.Find(&elements); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to retrieve elements"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Elements retrieved successfully", "data": elements})
}
func GetElement(c *gin.Context) {
    var element models.Element
    id := c.Param("id")
    if result := database.DB.First(&element, id); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to retrieve element"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Element retrieved successfully", "data": element})
}

func UpdateElement(c *gin.Context) {
    var element models.Element
    id := c.Param("id")

    if err := c.ShouldBindJSON(&element); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid data"})
        return
    }

    isValid, validationError := validateElement(&element)
    if !isValid {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": validationError})
        return
    }

    if result := database.DB.Model(&element).Where("id = ?", id).Updates(element); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to update element"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Element updated successfully", "data": element})
}

func DeleteElement(c *gin.Context) {
    id := c.Param("id")
    if result := database.DB.Delete(&models.Element{}, id); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to delete element"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Element deleted successfully"})
}