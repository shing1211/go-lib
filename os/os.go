package os

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func GetOSEnv(key, defaultValue string) string {
	log.SetFormatter(&log.JSONFormatter{})
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
