package main

import (
	"log"
	"portfolioBackend/database"
	"portfolioBackend/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	database.ConnectDB()
	router := gin.Default()

	// این خط حذف شده است: router.LoadHTMLGlob("templates/*")

	// تنظیمات کامل و صحیح CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// مسیرهای عمومی API
	api := router.Group("/api")
	{
		api.POST("/contact", handlers.CreateContactMessage)
		api.GET("/content", handlers.GetAllContent)
	}

	// گروه مسیرهای ادمین با احراز هویت
	adminRoutes := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "password123",
	}))
	{
		adminRoutes.GET("/projects", handlers.GetProjects)
		adminRoutes.POST("/projects", handlers.CreateProject)
		adminRoutes.PUT("/projects/:id", handlers.UpdateProject)
		adminRoutes.DELETE("/projects/:id", handlers.DeleteProject)
		adminRoutes.GET("/messages", handlers.GetContactMessages)
		adminRoutes.DELETE("/messages/:id", handlers.DeleteContactMessage)
		adminRoutes.PUT("/messages/:id/read", handlers.MarkAsRead) // <-- مسیر جدید
		adminRoutes.GET("/stats", handlers.GetDashboardStats)
	}

	log.Println("Server is running on port 8080")
	router.Run(":8080")
}
