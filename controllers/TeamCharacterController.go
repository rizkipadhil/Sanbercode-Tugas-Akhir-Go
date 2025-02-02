package controllers

import (
    "net/http"
    "tugas-akhir/database"
    "tugas-akhir/models"
    "github.com/gin-gonic/gin"
)

func validateTeamCharacter(teamCharacter models.TeamCharacter) []string {
    var errors []string
    if teamCharacter.TeamID == 0 {
        errors = append(errors, "Team ID is required")
    }
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

func CreateTeamCharacter(c *gin.Context) {
    var teamCharacter models.TeamCharacter
    if err := c.ShouldBindJSON(&teamCharacter); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid data"})
        return
    }

    var team models.Team
    if result := database.DB.First(&team, teamCharacter.TeamID); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Team not found"})
        return
    }

    userName, _ := c.Get("username")
    if team.CreatedBy != userName && userName != "superadmin" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "You are not authorized to update this team character"})
        return
    }

    teamCharacter.CreatedBy = userName.(string)

    if validationErrors := validateTeamCharacter(teamCharacter); len(validationErrors) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": validationErrors})
        return
    }

    if result := database.DB.Create(&teamCharacter); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Creation failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Team character created successfully"})
}

func GetTeamCharacter(c *gin.Context) {
    id := c.Param("id")
    var teamCharacter models.TeamCharacter
    if result := database.DB.Preload("Character").Preload("Artifact").First(&teamCharacter, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Team character not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "data": teamCharacter})
}

func GetTeamCharacters(c *gin.Context) {
    var teamCharacters []models.TeamCharacter
    userRole, _ := c.Get("role")
    username, _ := c.Get("username")

    if userRole == "superadmin" {
        database.DB.Preload("Character").Preload("Artifact").Find(&teamCharacters)
    } else {
        database.DB.Preload("Character").Preload("Artifact").Joins("JOIN teams ON teams.id = team_characters.team_id").Where("teams.created_by = ?", username).Find(&teamCharacters)
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "data": teamCharacters})
}

func UpdateTeamCharacter(c *gin.Context) {
    id := c.Param("id")

    var teamCharacterUpdate struct {
        TeamID      uint   `json:"team_id"`
        CharacterID uint   `json:"character_id"`
        ArtifactID  uint   `json:"artifact_id"`
        TypeSet     string `json:"type_set"`
        Mechanism   string `json:"mechanism"`
    }

    if err := c.ShouldBindJSON(&teamCharacterUpdate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid data"})
        return
    }

    var teamCharacter models.TeamCharacter
    if result := database.DB.First(&teamCharacter, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Team character not found"})
        return
    }

    var teamUpdate models.Team
    if result := database.DB.First(&teamUpdate, teamCharacterUpdate.TeamID); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Team not found"})
        return
    }

    var teamCurrent models.Team
    if result := database.DB.First(&teamCurrent, teamCharacter.TeamID); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Team not found"})
        return
    }

    userName, _ := c.Get("username")
    if (teamUpdate.CreatedBy != userName && userName != "superadmin") || (teamCurrent.CreatedBy != userName && userName != "superadmin") {
        c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "You are not authorized to update this team character"})
        return
    }

    if validationErrors := validateTeamCharacter(models.TeamCharacter{
        TeamID:      teamCharacterUpdate.TeamID,
        CharacterID: teamCharacterUpdate.CharacterID,
        ArtifactID:  teamCharacterUpdate.ArtifactID,
        TypeSet:     teamCharacterUpdate.TypeSet,
        Mechanism:   teamCharacterUpdate.Mechanism,
    }); len(validationErrors) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": validationErrors})
        return
    }

    teamCharacter.CharacterID = teamCharacterUpdate.CharacterID
    teamCharacter.ArtifactID = teamCharacterUpdate.ArtifactID
    teamCharacter.TypeSet = teamCharacterUpdate.TypeSet
    teamCharacter.Mechanism = teamCharacterUpdate.Mechanism
    teamCharacter.UpdatedBy = userName.(string)

    if result := database.DB.Save(&teamCharacter); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Update failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Team character updated successfully"})
}

func DeleteTeamCharacter(c *gin.Context) {
    id := c.Param("id")
    var teamCharacter models.TeamCharacter
    if result := database.DB.First(&teamCharacter, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Team character not found"})
        return
    }

    var team models.Team
    if result := database.DB.First(&team, teamCharacter.TeamID); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Team not found"})
        return
    }

    userName, _ := c.Get("username")
    if team.CreatedBy != userName && userName != "superadmin" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "You are not authorized to delete this team character"})
        return
    }

    if result := database.DB.Delete(&teamCharacter); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Delete failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Team character deleted successfully"})
}