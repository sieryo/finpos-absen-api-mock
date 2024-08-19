package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID            string         `json:"id" gorm:"type:char(36);primaryKey;default:(uuid())"`
	Name          string         `json:"name" gorm:"not null;size:100"`
	Username      string         `json:"username" gorm:"unique;not null;size:100"`
	Password      string         `json:"password" gorm:"not null;size:255"`
	Email         string         `json:"email" gorm:"not null;size:100"`
	ActivatedAt   sql.NullTime   `json:"activated_at"`
	FaceEmbedding []float64      `json:"face_embedding" gorm:"type:json"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Attendances   []Attendances  `json:"attendances" gorm:"foreignKey:UserID"`
}

type UserInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
