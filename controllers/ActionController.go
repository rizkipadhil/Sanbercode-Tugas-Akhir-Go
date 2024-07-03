package controllers

import (
    "net/http"
    "tugas-akhir/database"
    "tugas-akhir/models"
    "github.com/gin-gonic/gin"
)

func VerifyTeam(c *gin.Context) {
    id := c.Param("id")
    var team models.Team
    if result := database.DB.First(&team, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Team not found"})
        return
    }

    if team.VerifyStatus != "pending" {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Only teams with pending status can be verified"})
        return
    }

    var request struct {
        Status string `json:"status"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid data"})
        return
    }

    if request.Status != "diterima" && request.Status != "ditolak" {
        c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid status"})
        return
    }

    team.VerifyStatus = request.Status
    userName, _ := c.Get("username")
    team.VerifyBy = userName.(string)

    if result := database.DB.Save(&team); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to update status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Team status updated successfully"})
}