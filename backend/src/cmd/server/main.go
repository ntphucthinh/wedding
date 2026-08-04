package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"wedding/internal/di"
	"wedding/internal/interfaces/controller"
	"wedding/internal/pkg/app"
	"wedding/internal/usecases"
)

func main() {
	loadEnv()

	// Ưu tiên dùng DB_DSN nếu có, nếu không thì build từ từng biến lẻ
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		host := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		// Fallback mặc định nếu thiếu biến môi trường lẻ
		if user == "" {
			user = "root"
		}
		if pass == "" {
			pass = "root"
		}
		if host == "" {
			host = "wedding-db"
		}
		if dbPort == "" {
			dbPort = "3306"
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

	appHandler := di.InitializeHandler(db)

	authUsecase := usecases.NewAuthUsecase(nil)
	_ = authUsecase

	// Register routes once through the app handler.
	r := buildRouter(appHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func buildRouter(appHandler *controller.AppHandler) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	appHandler.RegisterRoutes(r)

	return r
}

func loadEnv() {
	if app.IsLocal() {
		// Nạp file .env.local tương tự như dự án store của bạn
		err := godotenv.Overload(".env.local")
		if err != nil {
			log.Println("Không thể nạp file .env.local, sử dụng biến môi trường hệ thống")
		}
	}
}
