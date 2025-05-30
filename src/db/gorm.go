package db

import (
	"fmt"

	"level-scale/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     uint16
	User     string
	Password string
	DbName   string
	Ssl      bool
}

var Db *gorm.DB

func Init(c Config) {
	sslMode := "disable"
	if c.Ssl {
		sslMode = "enable"
	}
	dsn := fmt.Sprintf("host=%s User=%s Password=%s DbName=%s port=%d sslMode=%s", c.Host, c.User, c.Password, c.DbName, c.Port, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatalw("Database initialization failed", "err", err)
	}
	Db = db
	logger.Log.Debugw("Database initialized")
}
