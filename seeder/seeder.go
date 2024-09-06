package seeder

import (
	"finpos-absen-api/config"
	"finpos-absen-api/internal/models"
)

func Init() {
	config.InitDatabase()
}

func SeedTipe() error {
	Init()
	types := []models.Tipe{
		{Tipe: "Kantor"},
		{Tipe: "WFH"},
		{Tipe: "Dinas Luar"},
		{Tipe: "Lembur"},
		{Tipe: "Piket"},
		{Tipe: "Masuk Kantor + WFH"},
		{Tipe: "WFH + Masuk Kantor"},
		{Tipe: "Sakit"},
		{Tipe: "Izin"},
		{Tipe: "Pengganti"},
	}

	for _, tipe := range types {
		if err := config.DB.FirstOrCreate(&tipe, models.Tipe{Tipe: tipe.Tipe}).Error; err != nil {
			return err
		}
	}

	return nil
}
