package env

import (
	"errors"
	"fmt"
	"os"
)

const (
	pgEnvHost     = "PG_HOST"
	pgEnvPort     = "PG_PORT"
	pgEnvDBName   = "POSTGRES_DB"
	pgEnvUser     = "POSTGRES_USER"
	pgEnvPassword = "POSTGRES_PASSWORD"
)

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func NewPGConfig() (PGConfig, error) {
	host := os.Getenv(pgEnvHost)
	if len(host) == 0 {
		return nil, errors.New("pg host not found")
	}
	port := os.Getenv(pgEnvPort)
	if len(port) == 0 {
		return nil, errors.New("pg port not found")
	}
	dbName := os.Getenv(pgEnvDBName)
	if len(dbName) == 0 {
		return nil, errors.New("pg dbName not found")
	}
	user := os.Getenv(pgEnvUser)
	if len(user) == 0 {
		return nil, errors.New("pg user not found")
	}
	pass := os.Getenv(pgEnvPassword)
	if len(pass) == 0 {
		return nil, errors.New("pg pass not found")
	}
	return &pgConfig{
		dsn: fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, dbName, user, pass),
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
