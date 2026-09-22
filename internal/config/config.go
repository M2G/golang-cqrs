package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Config struct {
	DatabaseURL string
	UploadDir   string
	StreamsDir  string
	HTTPAddr    string
	LogLevel    logrus.Level
}

func Load() (Config, error) {
	var missing []string

	require := func(key string) string {
		value := os.Getenv(key)
		if value == "" {
			missing = append(missing, key)
		}
		return value
	}

	cfg := Config{
		DatabaseURL: require("DATABASE_URL"),
		UploadDir:   require("UPLOAD_DIR"),
		StreamsDir:  require("STREAMS_DIR"),
		HTTPAddr:    getEnv("HTTP_ADDR", ":8181"),
	}

	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = logrus.InfoLevel
	}
	cfg.LogLevel = level

	if len(missing) > 0 {
		return Config{}, fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
