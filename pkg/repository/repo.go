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

func (r *partsRepository) GetParts() ([]parts.Part, error) {
	var partlist []parts.Part

	result, err := r.db.Query("SELECT `id`, `description`, `brand`, `color`, `price` from `computer_parts`")
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var part parts.Part
		err := result.Scan(&part.ID, &part.Description, &part.Brand, &part.Color, &part.NoTaxPrice)
		if err != nil {
			return nil, err
		}
		partlist = append(partlist, part)
	}
	return partlist, nil
}
