package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Define a custom type for FaceEmbedding
type Float64Slice []float64

// Implement the sql.Scanner interface
func (f *Float64Slice) Scan(value interface{}) error {
	if value == nil {
		*f = nil
		return nil
	}
	// Convert the []uint8 (JSON bytes) from the DB to []float64
	bytes, ok := value.([]uint8)
	if !ok {
		return errors.New("type assertion to []uint8 failed")
	}

	return json.Unmarshal(bytes, f)
}

// Implement the driver.Valuer interface
func (f Float64Slice) Value() (driver.Value, error) {
	if f == nil {
		return nil, nil
	}
	// Convert []float64 to JSON bytes for storing in the DB
	return json.Marshal(f)
}

type Users struct {
	ID            string         `json:"id" gorm:"type:char(36);primaryKey;default:(uuid())"`
	Name          string         `json:"name" gorm:"not null;size:100"`
	Username      string         `json:"username" gorm:"unique;not null;size:100"`
	Password      string         `json:"password" gorm:"not null;size:255"`
	Email         string         `json:"email" gorm:"not null;size:100"`
	ActivatedAt   sql.NullTime   `json:"activated_at"`
	FaceEmbedding Float64Slice   `json:"face_embedding" gorm:"type:json"`
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
