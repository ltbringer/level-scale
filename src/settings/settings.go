package settings

import (
	"level-scale/db"
	"level-scale/logger"
	"os"
	"strconv"
)

var (
	JWTSecret []byte
	DbConfig  *db.Config
)

func Init() {
	JWTSecret = []byte(getEnv("JWT_SECRET"))
	DbConfig = dbConfigFromEnv()
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		logger.Log.Fatalf("Missing required env var: %s", key)
	}
	return val
}

func parsePort(key string) uint16 {
	val := os.Getenv(key)
	if val == "" {
		logger.Log.Fatalf("Missing required env var: %s", key)
	}
	p, err := strconv.ParseUint(val, 10, 16)
	if err != nil {
		logger.Log.Fatalf("Invalid uint16 for %s: %v", key, err)
	}
	return uint16(p)
}

func stringToBool(value string) bool {
	b, err := strconv.ParseBool(value)
	if err != nil {
		logger.Log.Fatalf("Invalid value for %s: %v", value, err)
	}
	return b
}

func dbConfigFromEnv() *db.Config {
	c := &db.Config{
		Host:     getEnv("DB_HOST"),
		Port:     parsePort("DB_PORT"),
		User:     getEnv("DB_USER"),
		Password: getEnv("DB_PASS"),
		DbName:   getEnv("DB_NAME"),
		Ssl:      stringToBool(getEnv("DB_SSL")),
	}
	return c
}
