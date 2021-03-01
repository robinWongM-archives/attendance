package handlers

import (
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.GET("/api/auth", AuthStart)
	e.GET("/api/auth/callback", AuthCallback)

	e.GET("/api/users", GetUsers)
}
