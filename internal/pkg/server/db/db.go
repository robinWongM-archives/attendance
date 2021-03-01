package db

import (
	"github.com/robinWongM/attendance/internal/pkg/server/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dsn = "host=localhost user=ecnc password=z9yqKmO1CYGBG3abv23ZLtKrCdgbrYlL dbname=attendance port=5432 sslmode=disable TimeZone=Asia/Shanghai"
)

var db *gorm.DB

func Init() {
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to perform migration")
	}
}

func GetDB() *gorm.DB {
	return db
}
