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
)

func sendResponse(c *gin.Context, statusCode int, err bool, message string, data interface{}) {
    c.JSON(statusCode, gin.H{"error": err, "message": message, "data": data})
}

func ValidateUser(user *models.User) (bool, string) {
    if user.Username == "" {
        return false, "Username is required"
    }
    if user.Email == "" || !strings.Contains(user.Email, "@") {
        return false, "Valid email is required"
    }
    if user.Password == "" {
        return false, "Password is required"
    }
    return true, ""
}

func Register(c *gin.Context) {
    var user models.User
    var userInput struct {
        models.User
        PasswordConfirmation string `json:"password_confirmation"`
    }

    if err := c.ShouldBindJSON(&userInput); err != nil {
        sendResponse(c, http.StatusBadRequest, true, "Invalid data", nil)
        return
    }

    user = userInput.User

    isValid, validationError := ValidateUser(&user)
    if !isValid {
        sendResponse(c, http.StatusBadRequest, true, validationError, nil)
        return
    }

    if len(user.Password) < 6 {
        sendResponse(c, http.StatusBadRequest, true, "Password must be at least 6 characters long", nil)
        return
    }

    if user.Password != userInput.PasswordConfirmation {
        sendResponse(c, http.StatusBadRequest, true, "Password confirmation does not match", nil)
        return
    }

    var existingUser models.User
    if err := database.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error; err == nil {
        sendResponse(c, http.StatusBadRequest, true, "Username or email already registered", nil)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        sendResponse(c, http.StatusInternalServerError, true, "Failed to hash password", nil)
        return
    }
    user.Password = string(hashedPassword)
    user.Role = "member"

    if result := database.DB.Create(&user); result.Error != nil {
        sendResponse(c, http.StatusInternalServerError, true, "Registration failed", nil)
        return
    }

    sendResponse(c, http.StatusOK, false, "Registration successful", gin.H{"username": user.Username, "email": user.Email})
}

func Login(c *gin.Context) {
    var loginParams struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&loginParams); err != nil {
        sendResponse(c, http.StatusBadRequest, true, "Invalid request data", nil)
        return
    }

    if loginParams.Username == "" || loginParams.Password == "" {
        sendResponse(c, http.StatusBadRequest, true, "Username and password are required", nil)
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

    token, err := GenerateJWT(user.ID, user.Role, user.Username)
    if err != nil {
        sendResponse(c, http.StatusInternalServerError, true, "Failed to create token", nil)
        return
    }

    sendResponse(c, http.StatusOK, false, "Login successful", gin.H{"token": token})
}

func GenerateJWT(userID uint, role string, username string) (string, error) {
    var mySigningKey = []byte(os.Getenv("KEY_APP"))
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["user_id"] = userID
    claims["role"] = role
    claims["username"] = username
    claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

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
        newToken, err := GenerateJWT(uint(claims["user_id"].(float64)), claims["role"].(string), claims["username"].(string))
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