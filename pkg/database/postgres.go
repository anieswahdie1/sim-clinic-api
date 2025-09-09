package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sim-clinic-api/internal/config"
	"sim-clinic-api/internal/model"
)

func NewPostgresConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	// Configure GORM logger
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  getLogLevel(cfg.DBLogLevel),
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	logrus.Info("connected to database")
	return db, nil
}

func getLogLevel(level string) logger.LogLevel {
	switch level {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info
	}
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.BlacklistedToken{},
		&model.LayananTerapi{},
		&model.RiwayatPenyakit{},
		&model.TeknikTerapi{},
	)

	if err != nil {
		return err
	}

	// Seed initial roles
	return seedRoles(db)
}

func seedRoles(db *gorm.DB) error {
	roles := []model.Role{
		{Name: "super_admin", Description: "Super Administrator"},
		{Name: "admin", Description: "Administrator"},
		{Name: "user", Description: "Regular User"},
	}

	for _, role := range roles {
		var existingRole model.Role
		result := db.Where("name = ?", role.Name).First(&existingRole)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				if err := db.Create(&role).Error; err != nil {
					return err
				}
			} else {
				return result.Error
			}
		}
	}

	return nil
}
