package controllers

import (
    "net/http"
    "tugas-akhir/database"
    "tugas-akhir/models"
    "github.com/gin-gonic/gin"
)

func validateArtifact(artifact models.Artifact) []string {
    var errors []string
    if artifact.Name == "" {
        errors = append(errors, "Name is required")
    }
    if artifact.Description.Set2 == "" {
        errors = append(errors, "2-piece bonus description is required")
    }
    if artifact.Description.Set4 == "" {
        errors = append(errors, "4-piece bonus description is required")
    }
    if artifact.Rarity < 1 || artifact.Rarity > 5 {
        errors = append(errors, "Rarity must be between 1 and 5")
    }
    return errors
}

func CreateArtifact(c *gin.Context) {
    var artifact models.Artifact
    if err := c.ShouldBindJSON(&artifact); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid data"})
        return
    }

    if validationErrors := validateArtifact(artifact); len(validationErrors) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": validationErrors})
        return
    }

		userName, _ := c.Get("username")
		artifact.CreatedBy = userName.(string)

    if result := database.DB.Create(&artifact); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Creation failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Artifact created successfully"})
}

func GetArtifact(c *gin.Context) {
    id := c.Param("id")
    var artifact models.Artifact
    if result := database.DB.First(&artifact, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Artifact not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "data": artifact})
}

func GetArtifacts(c *gin.Context) {
    var artifacts []models.Artifact
    database.DB.Find(&artifacts)
    c.JSON(http.StatusOK, gin.H{"error": false, "data": artifacts})
}

func UpdateArtifact(c *gin.Context) {
    id := c.Param("id")
    var artifact models.Artifact
    if result := database.DB.First(&artifact, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Artifact not found"})
        return
    }

    if err := c.ShouldBindJSON(&artifact); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid data"})
        return
    }

    if validationErrors := validateArtifact(artifact); len(validationErrors) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": validationErrors})
        return
    }

		userName, _ := c.Get("username")
		artifact.UpdatedBy = userName.(string)

    if result := database.DB.Save(&artifact); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Update failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Artifact updated successfully"})
}

func DeleteArtifact(c *gin.Context) {
    id := c.Param("id")
    var artifact models.Artifact
    if result := database.DB.First(&artifact, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Artifact not found"})
        return
    }

    if result := database.DB.Delete(&artifact, id); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Delete failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Artifact deleted successfully"})
}