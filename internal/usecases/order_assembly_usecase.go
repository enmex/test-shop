package usecases

import (
	"test-shop/internal/dto"
	"test-shop/internal/interfaces"
	"test-shop/internal/models"
)

var _ interfaces.OrderAssemblyUsecase = (*OrderAssemblyUsecase)(nil)

type OrderAssemblyUsecase struct {
	orderRepo interfaces.OrderRepository
	productRepo interfaces.ProductRepository
}

func NewOrderAssemblyUsecase(
	orderRepo interfaces.OrderRepository,
	productRepo interfaces.ProductRepository,
) *OrderAssemblyUsecase {
	return &OrderAssemblyUsecase{orderRepo, productRepo}
}

func (u *OrderAssemblyUsecase) GetOrderAssembly(request *dto.GetOrderAssemblyRequest) (*dto.GetOrderAssemblyResponse, error) {
	orders, err := u.orderRepo.GetOrders(request.OrderIds)
	if err != nil {
		return nil, err
	}
	
	racksToOrders := make(map[string][]*models.Order)

	productToAdditionalRacks := make(map[int][]string, 0)
	for _, order := range orders {
		for _, product := range order.Products {
			additionalRacks, found := productToAdditionalRacks[product.ID]
			if !found {
				additionalRacks, err = u.productRepo.GetProductAdditionalRacks(product.ID)
				if err != nil {
					return nil, err
				}
			}
			product.AdditionalRacks = additionalRacks

			ordersInRack, found := racksToOrders[product.BasicRack]
			if !found {
				racksToOrders[product.BasicRack] = ordersInRack
			}
			ordersInRack = append(ordersInRack, order)
			racksToOrders[product.BasicRack] = ordersInRack
		}
	}
	return &dto.GetOrderAssemblyResponse{
		RacksToOrders: racksToOrders,
	}, nil
}