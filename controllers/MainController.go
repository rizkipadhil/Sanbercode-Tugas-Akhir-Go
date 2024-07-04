package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Greeting(c *gin.Context) {
    message := "Hallo, Selemat Datang Guest"

    // check if context has username
    if username, exists := c.Get("username"); exists {
        message = fmt.Sprintf("Hallo, Selamat Datang %s", username)
    }

    c.JSON(http.StatusOK, gin.H{"error": false, "message": message})
}