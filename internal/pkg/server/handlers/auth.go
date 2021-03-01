package handlers

import (
	"errors"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/robinWongM/attendance/internal/pkg/server/db"
	"github.com/robinWongM/attendance/internal/pkg/server/model"
	"gorm.io/gorm"
	"net/http"
)

func AuthStart(c echo.Context) error {
	// try to get the user without re-authenticating
	sess, err := session.Get("ecnc-attendance", c)
	if err != nil {
		return err
	}

	user, ok := sess.Values["user"].(goth.User)
	if ok {
		updateUserFromSSO(&user)
		return c.JSON(http.StatusOK, user)
	} else {
		gothic.BeginAuthHandler(c.Response(), c.Request())
		return nil
	}
}

func updateUserFromSSO(ssoUser *goth.User) error {
	db := db.GetDB()
	dbUser := model.User{}

	err := db.Where(&model.User{
		NetID: ssoUser.UserID,
	}).First(&dbUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		db.Create(&model.User{
			NetID: ssoUser.UserID,
			Name:  ssoUser.Name,
			Email: ssoUser.Email,
		})
	} else if err == nil {
		dbUser.Name = ssoUser.Name
		dbUser.Email = ssoUser.Email
		db.Save(&dbUser)
	} else {
		return err
	}

	return nil
}

func AuthCallback(c echo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return err
	}

	updateUserFromSSO(&user)

	sess, err := session.Get("ecnc-attendance", c)
	if err != nil {
		return err
	}

	sess.Values["user"] = user
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
