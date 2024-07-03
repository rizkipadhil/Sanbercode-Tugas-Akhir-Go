package controllers

import (
    "net/http"
    "tugas-akhir/database"
    "tugas-akhir/models"
    "github.com/gin-gonic/gin"
)

func validateTeam(team models.Team) []string {
    var errors []string
    if team.Name == "" {
        errors = append(errors, "Name is required")
    }
    if team.Description == "" {
        errors = append(errors, "Description is required")
    }
    return errors
}

func validateTeamCharacterTeam(teamCharacter models.TeamCharacter) []string {
    var errors []string
    if teamCharacter.CharacterID == 0 {
        errors = append(errors, "Character ID is required")
    }
    if teamCharacter.ArtifactID == 0 {
        errors = append(errors, "Artifact ID is required")
    }
    if teamCharacter.TypeSet == "" {
        errors = append(errors, "TypeSet is required")
    }
    if teamCharacter.Mechanism == "" {
        errors = append(errors, "Mechanism is required")
    }
    return errors
}

func CreateTeam(c *gin.Context) {
    var teamRequest struct {
        Name           string                    `json:"name"`
        Description    string                    `json:"description"`
        TeamCharacters []models.TeamCharacter    `json:"team_characters"`
    }

    if err := c.ShouldBindJSON(&teamRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid data"})
        return
    }

    team := models.Team{
        Name:        teamRequest.Name,
        Description: teamRequest.Description,
    }

    if validationErrors := validateTeam(team); len(validationErrors) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": validationErrors})
        return
    }

    userRole, _ := c.Get("role")
    userName, _ := c.Get("username")
    if userRole == "superadmin" {
        team.VerifyBy = userName.(string)
        team.VerifyStatus = "diterima"
    } else {
        team.VerifyStatus = "pending"
    }

    team.CreatedBy = userName.(string)

    tx := database.DB.Begin()
    if err := tx.Create(&team).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Creation failed"})
        return
    }

    for _, teamCharacter := range teamRequest.TeamCharacters {
        teamCharacter.TeamID = team.ID

        if validationErrors := validateTeamCharacterTeam(teamCharacter); len(validationErrors) > 0 {
            tx.Rollback()
            c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": validationErrors})
            return
        }

        if err := tx.Create(&teamCharacter).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to create team character"})
            return
        }
    }

    if err := tx.Commit().Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Transaction commit failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Team created successfully"})
}

func GetTeam(c *gin.Context) {
    id := c.Param("id")
    var team models.Team
    if result := database.DB.Preload("TeamCharacters.Character").Preload("TeamCharacters.Artifact").First(&team, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Team not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "data": team})
}

func GetTeams(c *gin.Context) {
    var teams []models.Team
    database.DB.Preload("TeamCharacters.Character").Preload("TeamCharacters.Artifact").Find(&teams)
    c.JSON(http.StatusOK, gin.H{"error": false, "data": teams})
}

func UpdateTeam(c *gin.Context) {
    id := c.Param("id")
    var team models.Team
    if result := database.DB.First(&team, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Team not found"})
        return
    }

    if err := c.ShouldBindJSON(&team); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid data"})
        return
    }

    if validationErrors := validateTeam(team); len(validationErrors) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": validationErrors})
        return
    }

    userName, _ := c.Get("username")
    team.UpdatedBy = userName.(string)

    if result := database.DB.Save(&team); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Update failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Team updated successfully"})
}

func DeleteTeam(c *gin.Context) {
    id := c.Param("id")
    var team models.Team
    if result := database.DB.First(&team, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Team not found"})
        return
    }

    if result := database.DB.Delete(&team, id); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Delete failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Team deleted successfully"})
}