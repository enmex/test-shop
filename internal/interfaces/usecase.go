package interfaces

import "test-shop/internal/dto"

type OrderAssemblyUsecase interface {
	GetOrderAssembly(request *dto.GetOrderAssemblyRequest) (*dto.GetOrderAssemblyResponse, error)
}
