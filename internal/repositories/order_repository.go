package repositories

import (
	"database/sql"
	"test-shop/internal/interfaces"
	"test-shop/internal/models"

	sq "github.com/Masterminds/squirrel"
)

var _ interfaces.OrderRepository = (*OrderRepository)(nil)

type OrderRepository struct {
	conn *sql.DB
}

func NewOrderRepository(conn *sql.DB) *OrderRepository {
	return &OrderRepository{
		conn,
	}
}

func (r *OrderRepository) GetOrders(orderIds []int) ([]*models.Order, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return nil, err
	}

	orders := make([]*models.Order, 0)
	for _, orderId := range orderIds {
		order := models.Order{
			ID: orderId,
		}

		rows, err := sq.Select("r.name", "p.id", "p.name", "quantity").
			From("products_orders po").
			LeftJoin("products p ON p.id = po.product_id").
			LeftJoin("products_racks pr ON pr.product_id = po.product_id").
			LeftJoin("racks r ON r.id = pr.rack_id").
			Where(sq.And{
				sq.Eq{"order_id": orderId},
				sq.Eq{"is_basic": true},
			}).
			RunWith(tx).
			PlaceholderFormat(sq.Dollar).
			Query()
		if err != nil {
			return nil, err
		}

		productsInOrder := make([]*models.ProductInOrder, 0)

		for rows.Next() {
			var productInOrder models.ProductInOrder

			if err := rows.Scan(&productInOrder.BasicRack, &productInOrder.ID, &productInOrder.Name, &productInOrder.Quantity); err != nil {
				if err := tx.Rollback(); err != nil {
					return nil, err
				}
				return nil, err
			}

			productsInOrder = append(productsInOrder, &productInOrder)
		}
		order.Products = productsInOrder

		rows.Close()
		
		orders = append(orders, &order)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return orders, nil
}
