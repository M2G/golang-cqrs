package config

import "os"

type Config struct {
	DatabaseURL string
	UploadDir   string
	StreamsDir  string
	HTTPAddr    string
	LogLevel    string
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


	//...
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
