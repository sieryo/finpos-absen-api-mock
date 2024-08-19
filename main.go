package main

import (
	"finpos-absen-api/config"
	"finpos-absen-api/internal/controllers"
	"finpos-absen-api/internal/middlewares"
	"finpos-absen-api/internal/models"
	"flag"
	"fmt"
	"net/http"
	"os"

	migrations "finpos-absen-api/migrate"

	"github.com/gin-gonic/gin"
)

func init() {
	config.InitDatabase()
}

func main() {
	r := gin.Default()

	migrate := flag.Bool("migrate", false, "Run migrations")
	flag.Parse()

	if *migrate {
		fmt.Println("Migrating...")
		migrations.Migrate()
		fmt.Println("Migrate completed.")
		os.Exit(0)
	}

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Todo List",
		})
	})

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", controllers.CreateUser)
		authRoutes.POST("/login", controllers.Login)
	}

	attendanceRoutes := r.Group("/attendance", middlewares.CheckAuth)
	{
		attendanceRoutes.POST("/clockin", controllers.CreateAttendance)
		attendanceRoutes.POST("/update_clockin", controllers.UpdateClockInAttendance)
		attendanceRoutes.POST("/clockout", controllers.UpdateClockOutAttendance)
	}

	r.GET("/profile", middlewares.CheckAuth, controllers.GetUserProfile)

	r.Run()
}

func testUser(user models.Users) {
	config.DB.Create(&user)
}
