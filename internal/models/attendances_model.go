package models

import (
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
	ID             string         `json:"id" gorm:"type:char(36);primaryKey;default:(uuid())"`
	UserID         string         `json:"user_id"`
	User           Users          `json:"users" gorm:"foreignKey:UserID"`
	Date           time.Time      `json:"date"`
	AttendanceType AttendanceType `json:"attendance_type"`
	ClockInWFH     *time.Time     `json:"clockin_wfh"`
	ClockOutWFH    *time.Time     `json:"clockout_wfh"`
	ClockInKantor  *time.Time     `json:"clockin_kantor"`
	ClockOutKantor *time.Time     `json:"clockout_kantor"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type AttendancesCreate struct {
	AttendanceType AttendanceType `json:"attendance_type"`
}
