package controllers

import (
	"finpos-absen-api/config"
	"finpos-absen-api/internal/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {

	var authInput models.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.Users
	result := config.DB.Where("email = ?", authInput.Email).First(&user)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password incorrect"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authInput.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password incorrect"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
