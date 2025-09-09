package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	customMiddleware "sim-clinic-api/internal/middleware"
	"sim-clinic-api/internal/service"
)

func SetupRoutes(e *echo.Echo, authService service.AuthService, userService service.UserService, masterService service.MasterDataService) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(LoggingMiddleware())
	e.Use(customMiddleware.AuthMiddleware(authService))

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})

	// Initialize handlers
	authHandler := NewAuthHandler(authService)
	userHandler := NewUserHandler(userService)
	masterHandler := NewMasterDataHandler(masterService)

	// API Group dengan prefix api
	api := e.Group("/api")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/logout", authHandler.Logout)
		}

		// Users Routes (protected)
		users := api.Group("/users")
		users.Use(customMiddleware.AuthMiddleware(authService))
		{
			users.GET("", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.PUT("/:id", userHandler.UpdateUser) // Tambahkan ini
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		master := api.Group("/master")
		master.Use(customMiddleware.AuthMiddleware(authService))
		{
			// Layanan Terapi
			layanan := master.Group("/layanan-terapi")
			{
				layanan.POST("", masterHandler.CreateLayananTerapi)
				layanan.GET("", masterHandler.GetAllLayananTerapi)
				layanan.GET("/:id", masterHandler.GetLayananTerapiByID)
				layanan.PUT("/:id", masterHandler.UpdateLayananTerapi)
				layanan.DELETE("/:id", masterHandler.DeleteLayananTerapi)
			}

			// Riwayat Penyakit
			riwayat := master.Group("/riwayat-penyakit")
			{
				riwayat.POST("", masterHandler.CreateRiwayatPenyakit)
				riwayat.GET("", masterHandler.GetAllRiwayatPenyakit)
				riwayat.GET("/:id", masterHandler.GetRiwayatPenyakitByID)
				riwayat.PUT("/:id", masterHandler.UpdateRiwayatPenyakit)
				riwayat.DELETE("/:id", masterHandler.DeleteRiwayatPenyakit)
			}

			// Teknik Terapi
			teknik := master.Group("/teknik-terapi")
			{
				teknik.POST("", masterHandler.CreateTeknikTerapi)
				teknik.GET("", masterHandler.GetAllTeknikTerapi)
				teknik.GET("/:id", masterHandler.GetTeknikTerapiByID)
				teknik.PUT("/:id", masterHandler.UpdateTeknikTerapi)
				teknik.DELETE("/:id", masterHandler.DeleteTeknikTerapi)
			}
		}
	}

}

func LoggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Before request
			logrus.WithFields(logrus.Fields{
				"method": c.Request().Method,
				"uri":    c.Request().URL.Path,
				"ip":     c.RealIP(),
			}).Info("Incoming request")

			err := next(c)

			// After request
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"method": c.Request().Method,
					"uri":    c.Request().URL.Path,
					"status": c.Response().Status,
					"error":  err.Error(),
				}).Error("Request failed")
			} else {
				logrus.WithFields(logrus.Fields{
					"method": c.Request().Method,
					"uri":    c.Request().URL.Path,
					"status": c.Response().Status,
				}).Info("Request completed")
			}

			return err
		}
	}
}
