package config

import (
	"flag"

	"github.com/joho/godotenv"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "deployments/env/.env", "path to config file; example: -config-path .env")
	flag.Parse()
}

func Load() error {
	err := godotenv.Load(configPath)
	if err != nil {
		return err
	}
	return nil
}

type LogConfig interface {
	Level() string
}

type GRPCConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}
