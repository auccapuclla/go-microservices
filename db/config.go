package db

import (
	"fmt"
	"os"
	"strconv"
)

type Config interface {
	GetDSN() string
	GetDBName() string
}

type config struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort int
	dbName string
	dsn    string
}

func NewConfig() Config {
	var cfg config
	cfg.dbUser = os.Getenv("DB_USER")
	cfg.dbPass = os.Getenv("DB_PASS")
	cfg.dbHost = os.Getenv("DB_HOST")
	cfg.dbName = os.Getenv("DB_NAME")
	var err error
	cfg.dbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	cfg.dbName = os.Getenv("DB_NAME")
	cfg.dsn = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s/?retryWrites=true&w=majority",
		cfg.dbUser, cfg.dbPass, cfg.dbHost, cfg.dbPort, cfg.dbName)
	cfg.dsn = os.Getenv("DB_DSN")
	return &cfg
}

func (c *config) GetDSN() string {
	return c.dsn
}

func (c *config) GetDBName() string {
	return c.dbName
}
