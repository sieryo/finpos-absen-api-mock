package config

import (
	"log"
	"time"

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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.Exec("CREATE DATABASE IF NOT EXISTS todo_app")

	dsn = "root:@tcp(127.0.0.1:3306)/todo_app?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the todo_app database: %v", err)
	}

	createTodoTable()
}

func createTodoTable() {
	DB.AutoMigrate(&Todo{})
}
