package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Address                    string `env:"SERVER_ADDRESS" envDefault:":8080"`
	DSN                        string `env:"DSN" envDefault:"postgresql://postgres:postgres@localhost:5432/s6er?sslmode=disable"`
	AccessTokenSecret          string `env:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret         string `env:"REFRESH_TOKEN_SECRET"`
	AccessTokenLiveTimeMinutes int    `env:"ACCESS_TOKEN_LIVE_TIME_MINUTES" envDefault:"30"`
	RefreshTokenLiveTimeDays   int    `env:"REFRESH_TOKEN_LIVE_TIME_DAYS" envDefault:"1"`
}

var cfg Config

func New() (*Config, error) {
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
