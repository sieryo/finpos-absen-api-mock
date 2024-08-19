package controllers

import (
	"finpos-absen-api/internal/models"
	"finpos-absen-api/internal/repositories"
	"finpos-absen-api/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAttendance(c *gin.Context) {
	user, exists := c.Get("currentUser")
	var userInput models.AttendancesCreate

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	currentUser, _ := user.(models.Users)

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !utils.IsValidAttendanceType(userInput.AttendanceType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid AttendanceType"})
		return
	}

	attendance, err := repositories.CreateAttendance(currentUser.ID, userInput.AttendanceType, time.Now(), nil)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": attendance})
}

func UpdateClockInAttendance(c *gin.Context) {
	user, exists := c.Get("currentUser")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	currentUser, _ := user.(models.Users)

	status, err := repositories.UpdateClockInAttendance(currentUser.ID, time.Now())

	if status != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": status})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdateClockOutAttendance(c *gin.Context) {
	user, exists := c.Get("currentUser")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	currentUser, _ := user.(models.Users)

	status, err := repositories.UpdateClockOutAttendance(currentUser.ID, time.Now())

	if status != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": status})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
