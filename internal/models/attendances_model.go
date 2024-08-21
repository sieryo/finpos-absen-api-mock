package models

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type AbsensiBase struct {
	ID            string         `json:"id" gorm:"type:char(36);primaryKey"`
	UserID        string         `json:"user_id"`
	User          Users          `json:"users" gorm:"foreignKey:UserID"`
	Tanggal       time.Time      `json:"tanggal" gorm:"type:date"`
	Tipe          uint64         `json:"tipe"`
	Clockin       *time.Time     `json:"clockin" gorm:"type:datetime(0)"`
	Clockout      *time.Time     `json:"clockout" gorm:"type:datetime(0)"`
	Foto          *string        `json:"foto" gorm:"type:varchar(100)"`
	Confidence    *float64       `json:"confidence" gorm:"type:double"`
	Emotion       *string        `json:"emotion" gorm:"type:varchar(20)"`
	FotoOut       *string        `json:"foto_out" gorm:"type:varchar(100)"`
	ConfidenceOut *float64       `json:"confidence_out" gorm:"type:double"`
	EmotionOut    *string        `json:"emotion_out" gorm:"type:varchar(20)"`
	Alasan        *string        `json:"alasan" gorm:"type:text"`
	Latitude      *string        `json:"latitude" gorm:"type:varchar(100)"`
	Longitude     *string        `json:"longitude" gorm:"type:varchar(100)"`
	LatitudeOut   *string        `json:"latitude_out" gorm:"type:varchar(100)"`
	LongitudeOut  *string        `json:"longitude_out" gorm:"type:varchar(100)"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Absensi struct {
	AbsensiBase
}

func (Absensi) TableName() string {
	return "absensi"
}

type AbsensiWFH struct {
	AbsensiBase
}

func (AbsensiWFH) TableName() string {
	return "absensi_wfh"
}

type AbsensiRequest struct {
	Tipe       uint64                `form:"tipe"`
	Foto       *multipart.FileHeader `form:"foto" binding:"required"`
	Confidence *string               `form:"confidence" binding:"required"`
	Latitude   *string               `form:"latitude"`
	Longitude  *string               `form:"longitude"`
}

type Tipe struct {
	ID        uint64         `json:"id" gorm:"type:bigint;primaryKey;autoIncrement"`
	Tipe      string         `json:"tipe"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Tipe) TableName() string {
	return "tipe"
}
