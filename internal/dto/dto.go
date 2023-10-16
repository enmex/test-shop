package dto

import "test-shop/internal/models"

type GetOrderAssemblyRequest struct {
	OrderIds []int
}

type GetOrderAssemblyResponse struct {
	RacksToOrders map[string][]*models.Order
}