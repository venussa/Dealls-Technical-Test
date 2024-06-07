package handlers

import (
  "net/http"
  "login-api/database"
  "login-api/models"
  "github.com/gin-gonic/gin"
  "golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
  Username string `json:"username" binding:"required"`
  Password string `json:"password" binding:"required"`
}

type LoginInput struct {
  Username string `json:"username" binding:"required"`
  Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
  var input RegisterInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
    return
  }

  user := models.User{Username: input.Username, Password: string(hashedPassword)}
  if err := database.DB.Create(&user).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func Login(c *gin.Context) {
  var input LoginInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  var user models.User
  if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    return
  }

  if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}