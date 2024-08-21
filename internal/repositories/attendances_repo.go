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

func HandleClockIn(userID string, tipe uint64, foto string, confidence float64, latitude string, longitude string) error {
	today := time.Now().Format("2006-01-02")
	now := time.Now().Truncate(time.Second)

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

		return nil
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
		return nil

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
}

func HandleClockOut(userID string, tipe uint64, foto string, confidence float64, latitude string, longitude string) error {
	today := time.Now().Format("2006-01-02")
	now := time.Now().Truncate(time.Second)

	switch tipe {
	case 1, 3, 4, 5, 10:
		// Masuk kantor, data ke absensi biasa
		var existingAttendance models.Absensi

		result := config.DB.Where("user_id = ? AND DATE(tanggal) = ?", userID, today).First(&existingAttendance)

		if result.Error != nil {
			return result.Error
		}

		existingAttendance.Clockout = &now
		existingAttendance.FotoOut = &foto
		existingAttendance.ConfidenceOut = &confidence
		existingAttendance.LatitudeOut = &latitude
		existingAttendance.LongitudeOut = &longitude

		if err := config.DB.Save(&existingAttendance).Error; err != nil {
			return err
		}

		return nil
	case 2:
		// WFH
		var existingAttendance models.AbsensiWFH

		result := config.DB.Where("user_id = ? AND DATE(tanggal) = ?", userID, today).First(&existingAttendance)

		if result.Error != nil {
			return result.Error
		}

		existingAttendance.Clockout = &now
		existingAttendance.FotoOut = &foto
		existingAttendance.ConfidenceOut = &confidence
		existingAttendance.LatitudeOut = &latitude
		existingAttendance.LongitudeOut = &longitude

		if err := config.DB.Save(&existingAttendance).Error; err != nil {
			return err
		}

		return nil

	case 6:
		// Masuk Kantor + WFH
		var absensi models.Absensi
		var absensiWFH models.AbsensiWFH

		result := config.DB.Where("user_id = ? AND DATE(tanggal) = ?", userID, today).First(&absensi)

		if result.Error != nil {
			return result.Error
		}

		result = config.DB.Where("user_id = ? AND DATE(tanggal) = ?", userID, today).First(&absensiWFH)

		if result.Error != nil {
			return result.Error
		}

		if absensi.Clockout == nil {
			absensi.Clockout = &now
			absensi.FotoOut = &foto
			absensi.ConfidenceOut = &confidence
			absensi.LatitudeOut = &latitude
			absensi.LongitudeOut = &longitude
			if err := config.DB.Save(&absensi).Error; err != nil {
				return err
			}

			return nil
		} else {
			absensiWFH.Clockout = &now
			absensiWFH.FotoOut = &foto
			absensiWFH.ConfidenceOut = &confidence
			absensiWFH.LatitudeOut = &latitude
			absensiWFH.LongitudeOut = &longitude

			if err := config.DB.Save(&absensiWFH).Error; err != nil {
				return err
			}

			return nil
		}
	case 7:
		// WFH + Masuk Kantor
		var absensi models.Absensi
		var absensiWFH models.AbsensiWFH

		result := config.DB.Where("user_id = ? AND DATE(tanggal) = ?", userID, today).First(&absensi)

		if result.Error != nil {
			return result.Error
		}

		result = config.DB.Where("user_id = ? AND DATE(tanggal) = ?", userID, today).First(&absensiWFH)

		if result.Error != nil {
			return result.Error
		}

		if absensiWFH.Clockout == nil {
			absensiWFH.Clockout = &now
			absensiWFH.FotoOut = &foto
			absensiWFH.ConfidenceOut = &confidence
			absensiWFH.LatitudeOut = &latitude
			absensiWFH.LongitudeOut = &longitude
			if err := config.DB.Save(&absensiWFH).Error; err != nil {
				return err
			}

			return nil
		} else {
			absensi.Clockout = &now
			absensi.FotoOut = &foto
			absensi.ConfidenceOut = &confidence
			absensi.LatitudeOut = &latitude
			absensi.LongitudeOut = &longitude

			if err := config.DB.Save(&absensi).Error; err != nil {
				return err
			}

			return nil
		}
	default:
		return fmt.Errorf("belum diimplementasikan")
	}
}
