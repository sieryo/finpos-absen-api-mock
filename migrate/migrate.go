package migrate

import (
	"finpos-absen-api/config"
	"finpos-absen-api/internal/models"
)

func Init() {
	config.InitDatabase()
}

func Migrate() {
	Init()
	config.DB.AutoMigrate(&models.Attendances{})
	config.DB.AutoMigrate(&models.Users{})

}
