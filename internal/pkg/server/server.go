package server

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/markbates/goth"
	"github.com/robinWongM/attendance/internal/pkg/server/db"
	"github.com/robinWongM/attendance/internal/pkg/server/handlers"
	"github.com/robinWongM/attendance/internal/pkg/sso"
	"os"
)

type Server struct {
	*echo.Echo
}

func initAuth(e *echo.Echo) {
	goth.UseProviders(
		sso.New(os.Getenv("ECNC_SSO_CLIENT_KEY"), os.Getenv("ECNC_SSO_CLIENT_SECRET"), "http://localhost:8080/api/auth/callback"),
	)
}

func NewServer() *Server {
	// Init DB
	db.Init()

	// Echo instance
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	// Middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewFilesystemStore("", []byte(os.Getenv("SESSION_SECRET")))))

	initAuth(e)
	handlers.InitRoute(e)

	return &Server{
		Echo: e,
	}
}
