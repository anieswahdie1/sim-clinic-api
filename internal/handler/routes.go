package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	customMiddleware "sim-clinic-api/internal/middleware"
	"sim-clinic-api/internal/service"
)

func SetupRoutes(e *echo.Echo, authService service.AuthService, userService service.UserService) {
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
		{
			users.GET("", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUserByID)
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
