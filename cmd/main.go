package main

import (
	"sim-clinic-api/internal/config"
	"sim-clinic-api/internal/handler"
	"sim-clinic-api/internal/repository"
	"sim-clinic-api/internal/service"
	"sim-clinic-api/pkg/database"
	logger "sim-clinic-api/pkg/log"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	// Setup logger
	logger.Init()

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
	// if err := database.AutoMigrate(db); err != nil {
	// 	logrus.Fatal("Error migrating database:", err)
	// }

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	masterDataRepo := repository.NewMasterDataRepository(db)
	customerRepo := repository.NewCustomerRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, roleRepo, tokenRepo, cfg.JWTSecret, cfg.JWTExpire)
	userService := service.NewUserService(userRepo)
	masterDataService := service.NewMasterDataService(masterDataRepo)
	customerService := service.NewCustomerService(customerRepo)

	// Setup routes
	handler.SetupRoutes(
		e,
		authService,
		userService,
		masterDataService,
		customerService,
	)

	// Start server
	logrus.Infof("Server starting on port %s", cfg.AppPort)
	if err := e.Start(":" + cfg.AppPort); err != nil {
		logrus.Fatal("Error starting server:", err)
	}
}
