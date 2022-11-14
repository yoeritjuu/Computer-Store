package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/yoeritjuu/Computer-Store/pkg/config"
)

func ConnectMySQL(cfg config.MySQLConfig) (*sql.DB, error) {
	ctx := context.Background()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database))
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed database connection because: %s", err.Error()))
	}

	db.SetConnMaxLifetime(cfg.ConnMaxLifeTime)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed ping because: %s", err.Error()))
	}

	log.Println("MySQL database connected!")

	var version string

	err = db.QueryRowContext(ctx, "SELECT VERSION()").Scan(&version)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed version query because: %s", err.Error()))
	}

	log.Printf("MySQL database version %s\n", version)
	return db, nil
}
