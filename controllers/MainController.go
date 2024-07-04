package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "tugas-akhir/models"
    "tugas-akhir/database"
)

func Greeting(c *gin.Context) {
    message := "Hallo, Selemat Datang Guest"

    // check if context has username
    if username, exists := c.Get("username"); exists {
        message = fmt.Sprintf("Hallo, Selamat Datang %s", username)
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": message})
}

func GetTeamsPublic(c *gin.Context) {
    var teams []models.Team

    verifyStatus := "diterima"

    database.DB.
        Preload("TeamCharacters.Character").
        Preload("TeamCharacters.Character.Element").
        Preload("TeamCharacters.Character.Weapon").
        Preload("TeamCharacters.Artifact").
        Where("verify_status = ?", verifyStatus).
        Find(&teams)
    
    c.JSON(http.StatusOK, gin.H{"error": false, "data": teams})
}