package footballdata

import (
	"fmt"
	"net/url"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	FootballDataBaseURL url.URL `envconfig:"FOOTBALL_DATA_BASE_URL" required:"true"`
	FootballDataAPIKey  string  `envconfig:"FOOTBALL_DATA_API_KEY" required:"true"`
}

func ReadConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("parsing config: %w", err)
	}

	return cfg, nil
}
