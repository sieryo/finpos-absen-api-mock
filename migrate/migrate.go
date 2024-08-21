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
	config.DB.AutoMigrate(&models.Absensi{})
	config.DB.AutoMigrate(&models.AbsensiWFH{})
	config.DB.AutoMigrate(&models.Tipe{})

}
