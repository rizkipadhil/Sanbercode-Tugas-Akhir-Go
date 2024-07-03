package auth

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "tugas-akhir/database"
    "tugas-akhir/models"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v4"
    "os"
    "time"
    "strings"
    "github.com/go-playground/validator/v10"
)

// Validator instance
var validate = validator.New()

// Helper function to send JSON responses
func sendResponse(c *gin.Context, statusCode int, err bool, message string, data interface{}) {
    c.JSON(statusCode, gin.H{"error": err, "message": message, "data": data})
}

// Register handles new user registration
func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        sendResponse(c, http.StatusBadRequest, true, "Invalid data", nil)
        return
    }

    // Validate input
    err := validate.Struct(user)
    if err != nil {
        sendResponse(c, http.StatusBadRequest, true, "Validation failed", gin.H{"errors": err.Error()})
        return
    }

    if user.Role != "member" {
        sendResponse(c, http.StatusBadRequest, true, "Only 'member' role can be registered", nil)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        sendResponse(c, http.StatusInternalServerError, true, "Password hashing failed", nil)
        return
    }
    user.Password = string(hashedPassword)
    user.Role = "member" // Set role to member by default

    if result := database.DB.Create(&user); result.Error != nil {
        sendResponse(c, http.StatusInternalServerError, true, "Registration failed", nil)
        return
    }

    sendResponse(c, http.StatusOK, false, "Registration successful", nil)
}

func Login(c *gin.Context) {
    var loginParams struct {
        Username string `json:"username" validate:"required"`
        Password string `json:"password" validate:"required"`
    }
    if err := c.ShouldBindJSON(&loginParams); err != nil {
        sendResponse(c, http.StatusBadRequest, true, "Invalid request data", nil)
        return
    }

    // Validate input
    err := validate.Struct(loginParams)
    if err != nil {
        sendResponse(c, http.StatusBadRequest, true, "Validation failed", gin.H{"errors": err.Error()})
        return
    }

    var user models.User
    if result := database.DB.Where("username = ?", loginParams.Username).First(&user); result.Error != nil {
        sendResponse(c, http.StatusBadRequest, true, "User not found", nil)
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginParams.Password)); err != nil {
        sendResponse(c, http.StatusBadRequest, true, "Invalid credentials", nil)
        return
    }

    token, err := GenerateJWT(user.ID, user.Role)
    if err != nil {
        sendResponse(c, http.StatusInternalServerError, true, "Failed to create token", nil)
        return
    }

    sendResponse(c, http.StatusOK, false, "Login successful", gin.H{"token": token})
}

func GenerateJWT(userID uint, role string) (string, error) {
    var mySigningKey = []byte(os.Getenv("KEY_APP"))
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["user_id"] = userID
    claims["role"] = role
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    tokenString, err := token.SignedString(mySigningKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

func RefreshToken(c *gin.Context) {
    oldToken := c.GetHeader("Authorization")
    tokenString := strings.TrimPrefix(oldToken, "Bearer ")

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("KEY_APP")), nil
    })

    if err != nil {
        sendResponse(c, http.StatusUnauthorized, true, "Failed to parse token", nil)
        return
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        newToken, err := GenerateJWT(uint(claims["user_id"].(float64)), claims["role"].(string))
        if err != nil {
            sendResponse(c, http.StatusInternalServerError, true, "Failed to create new token", nil)
            return
        }

        oldTokenEntry := models.OldToken{Token: tokenString, CreatedAt: time.Now()}
        database.DB.Create(&oldTokenEntry)

        sendResponse(c, http.StatusOK, false, "Token refreshed", gin.H{"token": newToken})
    } else {
        sendResponse(c, http.StatusUnauthorized, true, "Invalid token", nil)
    }
}

func Logout(c *gin.Context) {
    oldToken := c.GetHeader("Authorization")
    tokenString := strings.TrimPrefix(oldToken, "Bearer ")

    oldTokenEntry := models.OldToken{Token: tokenString, CreatedAt: time.Now()}
    database.DB.Create(&oldTokenEntry)

    sendResponse(c, http.StatusOK, false, "Successfully logged out", nil)
}

func UserAuth(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        sendResponse(c, http.StatusUnauthorized, true, "User not authenticated", nil)
        return
    }

    var user models.User
    if result := database.DB.First(&user, userID); result.Error != nil {
        sendResponse(c, http.StatusInternalServerError, true, "User not found", nil)
        return
    }

    sendResponse(c, http.StatusOK, false, "User authenticated", gin.H{"user": gin.H{
        "id": user.ID,
        "username": user.Username,
        "email": user.Email,
        "created_at": user.CreatedAt,
        "updated_at": user.UpdatedAt,
    }})
}