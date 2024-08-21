package routes

import (
	"finpos-absen-api/internal/controllers"
	"finpos-absen-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", controllers.CreateUser)
		authRoutes.POST("/login", controllers.Login)
	}
}

func AttendanceRoutes(r *gin.Engine) {
	attendanceRoutes := r.Group("/attendance", middlewares.CheckAuth)
	{
		attendanceRoutes.POST("/clockin", controllers.HandleClockIn)
		attendanceRoutes.POST("/clockout", controllers.HandleClockOut)
	}
}

func ProfileRoutes(r *gin.Engine) {
	profileRoutes := r.Group("/profile", middlewares.CheckAuth)
	{
		profileRoutes.GET("/", controllers.GetUserProfile)
	}
}
