package repositories

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
)

type ProductRepository struct {
	conn *sql.DB
}

func NewProductRepository(conn *sql.DB) *ProductRepository {
	return &ProductRepository{
		conn,
	}
}

func (r *ProductRepository) GetProductAdditionalRacks(productId int) ([]string, error) {
	additionalRacks := make([]string, 0)

	rows, err := sq.Select("r.name").
		From("products_racks").
		LeftJoin("racks r ON r.id = rack_id").
		Where(sq.And{
			sq.Eq{"is_basic": false},
			sq.Eq{"product_id": productId},
		}).
		RunWith(r.conn).
		PlaceholderFormat(sq.Dollar).
		Query()
	if err != nil {
		log.Println("here")
		return nil, err
	}
	for rows.Next() {
		var additionalRack string

		if err := rows.Scan(&additionalRack); err != nil {
			return nil, err
		}
		additionalRacks = append(additionalRacks, additionalRack)
	}
	return additionalRacks, nil
}
