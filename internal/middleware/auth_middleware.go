package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"sim-clinic-api/internal/service"
	"strings"
)

func AuthMiddleware(authService service.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Skip auth untuk public routes
			if isPublicRoute(c.Path()) {
				return next(c)
			}

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.ErrUnauthorized
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.ErrUnauthorized
			}

			tokenString := parts[1]

			// Validate token
			user, err := authService.ValidateToken(tokenString)
			if err != nil {
				logrus.Warnf("Invalid token: %v", err)
				return echo.ErrUnauthorized
			}

			// Set user info in context
			c.Set("userID", user.ID)
			c.Set("username", user.Username)
			c.Set("userRole", user.Role.Name)

			return next(c)
		}
	}
}

func isPublicRoute(path string) bool {
	publicRoutes := []string{
		"/api/auth/login",
		"/api/auth/register",
		"/swagger/",
		//"/health",
	}

	for _, route := range publicRoutes {
		if strings.HasPrefix(path, route) {
			return true
		}
	}
	return false
}
