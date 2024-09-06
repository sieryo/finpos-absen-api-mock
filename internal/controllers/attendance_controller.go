package controllers

import (
	"finpos-absen-api/internal/models"
	"finpos-absen-api/internal/repositories"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTodayAbsensi(c *gin.Context) {
	user, exists := c.Get("currentUser")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	currentUser, _ := user.(models.Users)

	absensiData, err := repositories.GetTodayAbsensi(currentUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve attendance data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": absensiData,
	})
}

func HandleClockIn(c *gin.Context) {

	user, exists := c.Get("currentUser")
	var input models.AbsensiRequest

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	currentUser, _ := user.(models.Users)

	if err := c.ShouldBind(&input); err != nil {
		fmt.Printf("Kesini! %s", err)
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

	if err := repositories.HandleClockIn(currentUser.ID, input.Tipe, fileName, confidence, *input.Latitude, *input.Longitude, *input.Alasan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Clock in berhasil"})
}

func HandleClockOut(c *gin.Context) {
	user, exists := c.Get("currentUser")
	var input models.AbsensiRequest

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

	if err := repositories.HandleClockOut(currentUser.ID, input.Tipe, fileName, confidence, *input.Latitude, *input.Longitude, *input.Alasan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Clock out berhasil"})
}
