package main

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/robinWongM/attendance/internal/pkg/sso"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func initAuth(e *echo.Echo) {
	goth.UseProviders(
		sso.New(os.Getenv("ECNC_SSO_CLIENT_KEY"), os.Getenv("ECNC_SSO_CLIENT_SECRET"), "http://localhost:8080/api/auth/callback"),
	)
	e.GET("/api/auth/callback", func(c echo.Context) error {
		user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
		if err != nil {
			e.Logger.Error(err)
		}

		sess, err := session.Get("ecnc-attendance", c)
		if err != nil {
			e.Logger.Error(err)
		}

		sess.Values["user"] = user
		sess.Save(c.Request(), c.Response())

		return c.JSON(http.StatusOK, user)
	})

	e.GET("/api/auth", func(c echo.Context) error {
		// try to get the user without re-authenticating
		sess, err := session.Get("ecnc-attendance", c)
		if err != nil {
			return err
		}

		user := sess.Values["user"]
		if user != nil {
			return c.JSON(http.StatusOK, user)
		} else {
			gothic.BeginAuthHandler(c.Response(), c.Request())
			return nil
		}
	})
}

func main() {
	// Echo instance
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	// Middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewFilesystemStore("", []byte(os.Getenv("ECNC_SESSION_KEY")))))

	initAuth(e)

	e.GET("/long", func(c echo.Context) error {
		time.Sleep(10 * time.Second)
		return c.JSON(http.StatusOK, "OK")
	})

	// Routes
	e.GET("/", hello)

	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	e.Logger.Info("Received interrupt signal, shutting down the server gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
