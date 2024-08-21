package repositories

import (
	"finpos-absen-api/config"
	"finpos-absen-api/internal/models"
)

func CreateUser(u *models.UserInput) (models.Users, error) {
	user := models.Users{
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}

	result := config.DB.Create(&user)

	if result.Error != nil {
		return models.Users{}, result.Error
	}

	return user, nil
}
