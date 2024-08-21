package controllers

import (
	"finpos-absen-api/internal/models"
	"finpos-absen-api/internal/repositories"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCurrentUserAttendances(c *gin.Context) {
	user, exists := c.Get("currentUser")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	currentUser, _ := user.(models.Users)

	attendances, err := repositories.GetCurrentUserAttendances(currentUser.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": attendances})
}

func GetCurrentDayUserAttendance(c *gin.Context) {
	user, exists := c.Get("currentUser")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	currentUser, _ := user.(models.Users)

	attendance, err := repositories.GetCurrentDayUserAttendance(currentUser.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": attendance})
}

func HandleClockIn(c *gin.Context) {
	user, exists := c.Get("currentUser")
	var input models.ClockinRequest

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	currentUser, _ := user.(models.Users)

	fmt.Println("User", currentUser.ID)

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	file, err := c.FormFile("foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate random filename
	fileName := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(file.Filename))

	// Define storage path
	storagePath := fmt.Sprintf("storage/image/%s", fileName)

	// Create directory if not exists
	if err := os.MkdirAll(filepath.Dir(storagePath), os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
		return
	}

	// Save the file
	if err := c.SaveUploadedFile(file, storagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	confidenceFloat, err := strconv.ParseFloat(*input.Confidence, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid confidence value"})
		return
	}

	confidence := float64(confidenceFloat)

	if err := repositories.HandleClockIn(currentUser.ID, input.Tipe, fileName, confidence, *input.Latitude, *input.Longitude); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Clock in berhasil"})
}

func UpdateClockInAttendance(c *gin.Context) {
	user, exists := c.Get("currentUser")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	currentUser, _ := user.(models.Users)

	status, err := repositories.UpdateCurrentUserClockInAttendance(currentUser.ID, time.Now())

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

	status, err := repositories.UpdateCurrentUserClockOutAttendance(currentUser.ID, time.Now())

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
