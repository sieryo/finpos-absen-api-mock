package main

import (
	"finpos-absen-api/config"
	"finpos-absen-api/internal/routes"
	"finpos-absen-api/seeder"
	"flag"
	"fmt"
	"os"

	migrations "finpos-absen-api/migrate"

	"github.com/gin-gonic/gin"
)

func init() {
	config.InitDatabase()
	config.InitEnv()
}

var r *gin.Engine

func main() {
	r = gin.Default()

	migrate := flag.Bool("migrate", false, "Run migrations")
	seed := flag.Bool("seed", false, "Run seeder")
	flag.Parse()

	if *migrate {
		fmt.Println("Migrating...")
		migrations.Migrate()
		fmt.Println("Migrate completed.")
		os.Exit(0)
	} else if *seed {
		fmt.Println("seeding...")
		seeder.SeedTipe()
		fmt.Println("seed completed.")
		os.Exit(0)
	}

	routes.AuthRoutes(r)
	routes.AttendanceRoutes(r)
	routes.ProfileRoutes(r)

	r.Run()
}
