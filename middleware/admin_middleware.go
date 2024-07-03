package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, exists := c.Get("role")
        if !exists || userRole != "superadmin" {
            c.JSON(http.StatusForbidden, gin.H{"error": true, "message": "You do not have access"})
            c.Abort()
            return
        }
        c.Next()
    }
}