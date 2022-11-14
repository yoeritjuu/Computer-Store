package repository

import (
	"database/sql"

	"github.com/yoeritjuu/Computer-Store/pkg/parts"
)

type partsRepository struct {
	db *sql.DB
}

func NewPartsRepository(sql *sql.DB) parts.PartsRepository {
	return &partsRepository{
		db: sql,
	}
}
