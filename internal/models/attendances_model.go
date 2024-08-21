package models

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type AttendanceType string

const (
	WFH        AttendanceType = "WFH"
	Kantor     AttendanceType = "Kantor"
	WFH_Kantor AttendanceType = "WFH_Kantor"
	Lembur     AttendanceType = "Lembur"
	Dinas_Luar AttendanceType = "Dinas_Luar"
)

type Attendances struct {
	ID                  string         `json:"id" gorm:"type:char(36);primaryKey;default:(uuid())"`
	UserID              string         `json:"user_id"`
	User                Users          `json:"users" gorm:"foreignKey:UserID"`
	Date                time.Time      `json:"date"`
	AttendanceType      AttendanceType `json:"attendance_type"`
	ClockInWFH          *time.Time     `json:"clockin_wfh"`
	ClockOutWFH         *time.Time     `json:"clockout_wfh"`
	ClockInKantor       *time.Time     `json:"clockin_kantor"`
	ClockOutKantor      *time.Time     `json:"clockout_kantor"`
	PhotoInWFH          *string        `json:"photo_in_wfh"`
	PhotoOutWFH         *string        `json:"photo_out_wfh"`
	PhotoInKantor       *string        `json:"photo_in_kantor"`
	PhotoOutKantor      *string        `json:"photo_out_kantor"`
	ConfidenceInWFH     *float32       `json:"confidence_in_wfh"`
	ConfidenceOutWFH    *float32       `json:"confidence_out_wfh"`
	ConfidenceInKantor  *float32       `json:"confidence_in_kantor"`
	ConfidenceOutKantor *float32       `json:"confidence_out_kantor"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

type AbsensiBase struct {
	ID            string         `json:"id" gorm:"type:char(36);primaryKey"`
	UserID        string         `json:"user_id"`
	User          Users          `json:"users" gorm:"foreignKey:UserID"`
	Tanggal       time.Time      `json:"tanggal" gorm:"type:date"`
	Tipe          uint64         `json:"tipe"`
	Clockin       *time.Time     `json:"clockin"`
	Clockout      *time.Time     `json:"clockout"`
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

type AttendancesCreate struct {
	AttendanceType AttendanceType `json:"attendance_type" binding:"required"`
	Confidence     string         `json:"confidence" binding:"required"`
}

// Struct untuk edit attendance
type AttendanceEditInput struct {
	ID             string         `json:"id" binding:"required"`
	Date           *time.Time     `json:"date"`
	AttendanceType AttendanceType `json:"attendance_type" binding:"required"`
	ClockInWFH     *time.Time     `json:"clockin_wfh"`
	ClockOutWFH    *time.Time     `json:"clockout_wfh"`
	ClockInKantor  *time.Time     `json:"clockin_kantor"`
	ClockOutKantor *time.Time     `json:"clockout_kantor"`
}

type ClockinRequest struct {
	Tipe       uint64                `form:"tipe"`
	Foto       *multipart.FileHeader `form:"foto" binding:"required"`
	Confidence *string               `form:"confidence" binding:"required"`
	Latitude   *string               `form:"latitude"`
	Longitude  *string               `form:"longitude"`
}

type ClockoutRequest struct {
	FotoOut       *multipart.FileHeader `form:"foto_out" binding:"required"`
	ConfidenceOut *string               `form:"confidence_out" binding:"required"`
	LatitudeOut   *string               `form:"latitude_out"`
	LongitudeOut  *string               `form:"longitude_out"`
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
