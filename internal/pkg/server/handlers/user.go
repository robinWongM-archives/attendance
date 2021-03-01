package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/robinWongM/attendance/internal/pkg/server/db"
	"github.com/robinWongM/attendance/internal/pkg/server/model"
	"net/http"
)

func GetUsers(c echo.Context) error {
	db := db.GetDB()
	var users []model.User
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}
