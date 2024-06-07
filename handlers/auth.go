package handlers

import (
	http "net/http"
	mysql "github.com/venussa/Dealls-Technical-Test/database/driver"
	models "github.com/venussa/Dealls-Technical-Test/models"
	gin "github.com/gin-gonic/gin"
	bcrypt "golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Email string `json:"email" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	ProfilePicture string `json:"profile_picture" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email string `json:"email" binding:"required"`
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

	user := models.User{Email: input.Email,
		FullName: input.FullName,
		ProfilePicture: input.ProfilePicture,
		Password: string(hashedPassword)}

	if err := mysql.ConnectDatabase().Create(&user).Error; err != nil {
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
	if err := mysql.ConnectDatabase().Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}