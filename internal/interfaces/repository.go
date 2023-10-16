package interfaces

import "test-shop/internal/models"

type OrderRepository interface {
	GetOrders(orderIds []int) ([]*models.Order, error)
}

type ProductRepository interface {
	GetProductAdditionalRacks(productId int) ([]string, error)
}