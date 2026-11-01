package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	NaraAPIKey string
	NaraModel  string
}

// Load reads a project .env file when present. Existing environment variables take precedence.
func Load() (Config, error) {
	for _, path := range []string{".env", "../.env"} {
		if _, err := os.Stat(path); err == nil {
			if err := godotenv.Load(path); err != nil {
				return Config{}, fmt.Errorf("load %s: %w", path, err)
			}
			break
		} else if !os.IsNotExist(err) {
			return Config{}, fmt.Errorf("read %s: %w", path, err)
		}
	}
	cfg := Config{Port: strings.TrimSpace(os.Getenv("PORT")), NaraAPIKey: strings.TrimSpace(os.Getenv("NARA_API_KEY")), NaraModel: strings.TrimSpace(os.Getenv("NARA_MODEL"))}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.NaraModel == "" {
		cfg.NaraModel = "deepseek-v4.1-flash"
	}
	if cfg.NaraAPIKey == "" {
		return Config{}, fmt.Errorf("NARA_API_KEY is required; set it in the project .env or environment")
	}
	return cfg, nil
}
