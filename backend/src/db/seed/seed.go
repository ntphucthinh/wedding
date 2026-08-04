package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"wedding/db/seed/seeder"
)

func main() {
	godotenv.Load(".env.local")

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		user := os.Getenv("DB_USERNAME")
		pass := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_DATABASE")

		if user == "" {
			user = "root"
		}
		if pass == "" {
			pass = "root"
		}
		if host == "" {
			host = "localhost"
		}
		if dbPort == "" {
			dbPort = "3309"
		}
		if dbName == "" {
			dbName = "wedding_db"
		}

		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, dbPort, dbName)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := seeder.SeedAdminUser(db); err != nil {
		log.Fatalf("seed failed: %v", err)
	}

	fmt.Println("Seeding completed successfully.")
}
