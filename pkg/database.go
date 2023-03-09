package pkg

import (
	"github.com/devgoorita/golang-basecode/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "host=" + GodotEnv("POSTGRES_HOST") + " user=" + GodotEnv("POSTGRES_USER") + " password=" + GodotEnv("POSTGRES_PASSWORD") + " dbname=" + GodotEnv("POSTGRES_DB") + " port=" + GodotEnv("POSTGRES_PORT") + " sslmode=disable TimeZone=Asia/Jakarta search_path=" + GodotEnv("POSTGRES_SCHEMA")
	db, err := gorm.Open(
		postgres.Open(dsn), &gorm.Config{})

	// Migrate the schema
	migrateDatabase(db)

	if err != nil {
		defer logrus.Info("Database connection failed")
		logrus.Fatal(err)
		return nil
	}
	return db
}

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.User{})
}
