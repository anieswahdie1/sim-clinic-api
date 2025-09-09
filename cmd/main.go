package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
	"sim-clinic-api/internal/config"
	"sim-clinic-api/internal/handler"
	"sim-clinic-api/internal/repository"
	"sim-clinic-api/internal/service"
	"sim-clinic-api/pkg/database"
)

// @title SIM Clinic API
// @version 1.0
// @description API for SIM Clinic Application
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@simclinic.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Setup logger
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatal("Error loading config:", err)
	}

	// Initialize Echo
	e := echo.New()

	// Initialize database
	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		logrus.Fatal("Error connecting to database:", err)
	}

	// Auto migrate
	if err := database.AutoMigrate(db); err != nil {
		logrus.Fatal("Error migrating database:", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	masterDataRepo := repository.NewMasterDataRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, roleRepo, tokenRepo, cfg.JWTSecret, cfg.JWTExpire)
	userService := service.NewUserService(userRepo)
	masterDataService := service.NewMasterDataService(masterDataRepo)

	// Setup routes
	handler.SetupRoutes(e, authService, userService, masterDataService)

	// Start server
	logrus.Infof("Server starting on port %s", cfg.AppPort)
	if err := e.Start(":" + cfg.AppPort); err != nil {
		logrus.Fatal("Error starting server:", err)
	}
}
