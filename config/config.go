package config

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var DB *gorm.DB

func InitDatabase() {

	dsn := "root:@tcp(127.0.0.1:3306)/"
	dbName := "finpos_absen"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)

	db.Exec(query)

	dsn = fmt.Sprintf("root:@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=true&loc=Local", dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the todo_app database: %v", err)
	}
}

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
