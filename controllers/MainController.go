package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Greeting(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"error": false, "message": "Hallo, Selamat Datang"})
}