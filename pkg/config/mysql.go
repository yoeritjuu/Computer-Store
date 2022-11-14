package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type MySQLConfig struct {
	Username        string        `env:"MYSQL_USERNAME"`
	Password        string        `env:"MYSQL_PASSWORD"`
	Host            string        `env:"MYSQL_HOST"`
	Port            int           `env:"MYSQL_PORT"`
	Database        string        `env:"MYSQL_DATABASE"`
	ConnMaxLifeTime time.Duration `env:"CONN_MAX_LIFE_TIME"`
	MaxIdleConns    int           `env:"MAX_IDLE_CONNS"`
	MaxOpenConns    int           `env:"MAX_OPEN_CONNS"`
}

// Gets all values from the environment.
func LoadSqlConfig() MySQLConfig {
	connMaxLifeTime, err := time.ParseDuration(os.Getenv("CONN_MAX_LIFE_TIME"))
	if err != nil {
		log.Fatalf("cant convert maxConnLifeTime env")
	}
	MaxIdleConns, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))
	if err != nil {
		log.Fatalf("cant convert maxIdleConns env")
	}
	MaxOpenConns, err := strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))
	if err != nil {
		log.Fatalf("cant convert maxOpenConns env")
	}

	port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		log.Fatalf("cant convert port env")
	}
	return MySQLConfig{
		Username:        os.Getenv("MYSQL_USERNAME"),
		Password:        os.Getenv("MYSQL_PASSWORD"),
		Host:            os.Getenv("MYSQL_HOST"),
		Port:            port,
		Database:        os.Getenv("MYSQL_DATABASE"),
		ConnMaxLifeTime: connMaxLifeTime,
		MaxIdleConns:    MaxIdleConns,
		MaxOpenConns:    MaxOpenConns,
	}
}
