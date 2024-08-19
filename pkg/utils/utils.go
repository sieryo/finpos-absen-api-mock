package utils

import (
	"finpos-absen-api/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pasword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pasword), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

var validAttendanceTypes = map[models.AttendanceType]bool{
	models.WFH:        true,
	models.Kantor:     true,
	models.WFH_Kantor: true,
	models.Lembur:     true,
	models.Dinas_Luar: true,
}

func IsValidAttendanceType(at models.AttendanceType) bool {
	_, exists := validAttendanceTypes[at]
	return exists
}
