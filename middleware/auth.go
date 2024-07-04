package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/golang-jwt/jwt/v4"
    "os"
    "strings"
    "tugas-akhir/database"
    "tugas-akhir/models"
    "time"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Authorization header required"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("KEY_APP")), nil
        })

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            exp := int64(claims["exp"].(float64))
            if time.Now().Unix() > exp {
                oldToken := models.OldToken{Token: tokenString, CreatedAt: time.Now()}
                database.DB.Create(&oldToken)
                c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Token expired"})
                c.Abort()
                return
            }

            var oldTokens []models.OldToken
            database.DB.Where("token = ?", tokenString).Find(&oldTokens)
            if len(oldTokens) > 0 {
                c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Invalid token"})
                c.Abort()
                return
            }

            c.Set("user_id", claims["user_id"])
            c.Set("role", claims["role"])
            c.Set("username", claims["username"])
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Invalid token", "details": err.Error()})
            c.Abort()
            return
        }

        c.Next()
    }
}