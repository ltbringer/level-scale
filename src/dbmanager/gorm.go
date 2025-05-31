package dbmanager

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

func CreateDsn(c *Config) string {
	sslMode := "disable"
	if c.Ssl {
		sslMode = "enable"
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", c.Host, c.User, c.Password, c.DbName, c.Port, sslMode)
}

func Init(c *Config) {
	dsn := CreateDsn(c)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatalw("Database initialization failed", "err", err)
	}
	Db = db
	logger.Log.Debugw("Database initialized")
}
