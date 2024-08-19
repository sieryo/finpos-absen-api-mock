package repositories

import (
	"errors"
	"finpos-absen-api/config"
	"finpos-absen-api/internal/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func GetAttendanceById(id string) (models.Attendances, error) {
	var attendance models.Attendances

	result := config.DB.First(&attendance, id)

	if result.Error != nil {
		return models.Attendances{}, result.Error
	}

	return attendance, nil
}

func CreateAttendance(userID string, at models.AttendanceType, clockIn time.Time, clockOut *time.Time) (models.Attendances, error) {
	var existingAttendance models.Attendances

	result := config.DB.Where("user_id = ? AND DATE(date) = ?", userID, time.Now().Format("2006-01-02")).First(&existingAttendance)

	if result.Error == nil {
		return models.Attendances{}, fmt.Errorf("attendance already recorded for today")
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Attendances{}, fmt.Errorf("attendance already recorded for today")
	}

	attendance := models.Attendances{
		UserID:         userID,
		Date:           time.Now(),
		AttendanceType: at,
	}

	switch at {
	case models.WFH, models.WFH_Kantor:
		attendance.ClockInWFH = &clockIn
		if clockOut != nil {
			attendance.ClockOutWFH = clockOut
		}
	case models.Kantor, models.Lembur, models.Dinas_Luar:
		attendance.ClockInKantor = &clockIn
		if clockOut != nil {
			attendance.ClockOutKantor = clockOut
		}
	}

	if err := config.DB.Create(&attendance).Error; err != nil {
		return models.Attendances{}, err
	}

	return attendance, nil
}

func UpdateClockInAttendance(userID string, clockIn time.Time) (string, error) {
	var attendance models.Attendances

	result := config.DB.Where("user_id = ? AND DATE(date) = ?", userID, time.Now().Format("2006-01-02")).First(&attendance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("attendance record not found for today")
		}
		return "", result.Error
	}

	if attendance.ClockInKantor != nil {
		return "Attendance already recorded, cannot update clock-in time", nil
	}

	result = config.DB.Model(&attendance).Update("clock_in_kantor", &clockIn)
	if result.Error != nil {
		return "", result.Error
	}

	return "", nil
}

func UpdateClockOutAttendance(userID string, clockOut time.Time) (string, error) {
	var attendance models.Attendances

	result := config.DB.Where("user_id = ? AND DATE(date) = ?", userID, time.Now().Format("2006-01-02")).First(&attendance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("attendance record not found for today")
		}
		return "", result.Error
	}

	if attendance.ClockOutKantor != nil || attendance.ClockOutWFH != nil {
		return "Attendance already recorded, cannot update clock-out time", nil

	}

	switch attendance.AttendanceType {
	case models.WFH, models.WFH_Kantor:
		if attendance.ClockOutWFH != nil {
			attendance.ClockOutKantor = &clockOut
		} else {
			attendance.ClockOutWFH = &clockOut
		}
	default:
		attendance.ClockOutKantor = &clockOut
	}

	config.DB.Save(&attendance)

	return "", nil
}
