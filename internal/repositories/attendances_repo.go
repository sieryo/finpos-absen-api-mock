package repositories

import (
	"errors"
	"finpos-absen-api/config"
	"finpos-absen-api/internal/models"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetCurrentUserAttendances(ID string) ([]models.Attendances, error) {
	var attendances []models.Attendances

	result := config.DB.Where("user_id = ?", ID).Find(&attendances)

	if result.Error != nil {
		return []models.Attendances{}, result.Error
	}

	return attendances, nil
}

func GetCurrentDayUserAttendance(userID string) (models.Attendances, error) {
	var attendance models.Attendances

	result := config.DB.Where("user_id = ? AND DATE(date) = ?", userID, time.Now().Format("2006-01-02")).First(&attendance)

	if result.Error != nil {
		return models.Attendances{}, result.Error
	}

	return attendance, nil
}

func GetAttendanceById(id string) (models.Attendances, error) {
	var attendance models.Attendances

	result := config.DB.First(&attendance, id)

	if result.Error != nil {
		return models.Attendances{}, result.Error
	}

	return attendance, nil
}

func CreateAttendance(userID string, at models.AttendanceType, clockIn time.Time, photoInPath *string, confidence *float32) (models.Attendances, error) {
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

	case models.Kantor, models.Lembur, models.Dinas_Luar:
		attendance.ClockInKantor = &clockIn
	}

	if err := config.DB.Create(&attendance).Error; err != nil {
		return models.Attendances{}, err
	}

	return attendance, nil
}

func UpdateCurrentUserClockInAttendance(userID string, clockIn time.Time) (string, error) {
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

func UpdateCurrentUserClockOutAttendance(userID string, clockOut time.Time) (string, error) {
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

func EditAttendance(a models.AttendanceEditInput) (models.Attendances, error) {
	var attendance models.Attendances

	result := config.DB.First(&attendance, a.ID)
	if result.Error != nil {
		return models.Attendances{}, result.Error
	}

	attendance.Date = *a.Date
	attendance.AttendanceType = a.AttendanceType
	attendance.ClockInWFH = a.ClockInWFH
	attendance.ClockOutWFH = a.ClockOutWFH
	attendance.ClockInKantor = a.ClockInKantor
	attendance.ClockOutKantor = a.ClockOutKantor

	saveResult := config.DB.Save(&attendance)

	if saveResult.Error != nil {
		return models.Attendances{}, saveResult.Error
	}

	return attendance, nil
}

func DeleteAttendance(ID string) (bool, error) {

	result := config.DB.Delete(&models.Attendances{}, ID)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func HandleClockIn(userID string, tipe uint64, foto string, confidence float64, latitude string, longitude string) error {
	today := time.Now().Format("2006-01-02")
	now := time.Now().Truncate(time.Millisecond)

	switch tipe {
	case 1, 3, 4, 5, 10:
		// Masuk kantor, data ke absensi biasa
		var existingAttendance models.Absensi

		result := config.DB.Where("user_id = ? AND DATE(tanggal) = ?", userID, today).First(&existingAttendance)

		if result.Error == nil {
			return fmt.Errorf("attendance already recorded for today")
		} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("attendance already recorded for today")
		}

		absensi := models.Absensi{
			AbsensiBase: models.AbsensiBase{
				ID:         uuid.New().String(),
				UserID:     userID,
				Tanggal:    now,
				Tipe:       tipe,
				Clockin:    &now,
				Foto:       &foto,
				Confidence: &confidence,
				Latitude:   &latitude,
				Longitude:  &longitude,
			},
		}
		if err := config.DB.Create(&absensi).Error; err != nil {
			return err
		}
	case 2:
		// WFH
		var existingAttendance models.AbsensiWFH

		result := config.DB.Where("user_id = ? AND DATE(tanggal) = ?", userID, today).First(&existingAttendance)

		if result.Error == nil {
			return fmt.Errorf("attendance already recorded for today")
		} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("attendance already recorded for today")
		}

		absensiWFH := models.AbsensiWFH{
			AbsensiBase: models.AbsensiBase{
				ID:         uuid.New().String(),
				UserID:     userID,
				Tanggal:    now,
				Tipe:       tipe,
				Clockin:    &now,
				Foto:       &foto,
				Confidence: &confidence,
				Latitude:   &latitude,
				Longitude:  &longitude,
			},
		}
		if err := config.DB.Create(&absensiWFH).Error; err != nil {
			return err
		}
	case 6:
		// Masuk Kantor + WFH
		var absensiWFH models.AbsensiWFH

		result := config.DB.Where("user_id = ? AND DATE(tanggal) = ?", userID, today).First(&absensiWFH)

		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return result.Error
		}

		if absensiWFH.Clockin != nil {
			return fmt.Errorf("attendance already recorded for today")
		}

		if result.Error == gorm.ErrRecordNotFound {
			absensi := models.Absensi{
				AbsensiBase: models.AbsensiBase{
					ID:         uuid.New().String(),
					UserID:     userID,
					Tanggal:    now,
					Tipe:       tipe,
					Clockin:    &now,
					Foto:       &foto,
					Confidence: &confidence,
					Latitude:   &latitude,
					Longitude:  &longitude,
				},
			}
			config.DB.Create(&absensi)
			absensiWFH := models.AbsensiWFH{
				AbsensiBase: models.AbsensiBase{
					ID:      uuid.New().String(),
					UserID:  userID,
					Tanggal: now,
					Tipe:    tipe,
				},
			}
			if err := config.DB.Create(&absensiWFH).Error; err != nil {
				return err
			}

			return nil
		}
		absensiWFH.Clockin = &now
		absensiWFH.Foto = &foto
		absensiWFH.Confidence = &confidence
		absensiWFH.Latitude = &latitude
		absensiWFH.Longitude = &longitude

		return config.DB.Save(&absensiWFH).Error
	case 7:
		// WFH + Masuk Kantor
		var absensi models.Absensi

		result := config.DB.Where("user_id = ? AND DATE(tanggal) = ?", userID, today).First(&absensi)

		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return result.Error
		}

		if absensi.Clockin != nil {
			return fmt.Errorf("attendance already recorded for today")
		}

		if result.Error == gorm.ErrRecordNotFound {
			absensi := models.Absensi{
				AbsensiBase: models.AbsensiBase{
					ID:      uuid.New().String(),
					UserID:  userID,
					Tanggal: now,
					Tipe:    tipe,
				},
			}
			config.DB.Create(&absensi)
			absensiWFH := models.AbsensiWFH{
				AbsensiBase: models.AbsensiBase{
					ID:         uuid.New().String(),
					UserID:     userID,
					Tanggal:    now,
					Tipe:       tipe,
					Clockin:    &now,
					Foto:       &foto,
					Confidence: &confidence,
					Latitude:   &latitude,
					Longitude:  &longitude,
				},
			}
			if err := config.DB.Create(&absensiWFH).Error; err != nil {
				return err
			}

			return nil
		}
		absensi.Clockin = &now
		absensi.Foto = &foto
		absensi.Confidence = &confidence
		absensi.Latitude = &latitude
		absensi.Longitude = &longitude

		return config.DB.Save(&absensi).Error
	default:
		return fmt.Errorf("belum diimplementasikan")
	}
	return fmt.Errorf("belum diimplementasikan")

}
