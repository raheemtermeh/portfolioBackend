package database

import (
	"fmt"
	"log"
	"os"
	"portfolioBackend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB به دیتابیس PostgreSQL متصل می‌شود
func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
	}

	log.Println("Successfully connected to the database")

	// ساخت جدول به صورت خودکار
	DB.AutoMigrate(
        &models.ContactMessage{}, 
        &models.Skill{}, 
        &models.Project{}, 
        &models.TimelineEvent{},
        &models.SiteConfig{},
    )
	log.Println("Database migrated")
}