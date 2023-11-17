package usecase

import (
	"freigth_service/internal/application/dto"
	"freigth_service/internal/domain/entity"
	"freigth_service/internal/domain/gateway"
)

type CreateRouteUseCase struct {
	RouteRepository gateway.RouteRepositoryInterface
	Freight         *entity.Freight
}

func NewCreateRouteUseCase(repo gateway.RouteRepositoryInterface,freight *entity.Freight) *CreateRouteUseCase{
	return &CreateRouteUseCase{
		RouteRepository: repo,
		Freight: freight,
	}
}

func(usecase *CreateRouteUseCase) Execute(input dto.InputCreateRouteDto) (*dto.OutputCreateRouteDto,error) {
	route := entity.NewRoute(input.ID,input.Name,input.Distance)
	usecase.Freight.Calculate(route)
	err := usecase.RouteRepository.Create(route)
	if err != nil {
		return nil,err
	}
	return &dto.OutputCreateRouteDto{
		ID: route.ID,
		Name: route.Name,
		Distance: route.Distance,
		Status: route.Status,
		FreightPrice: route.FreightPrice,
	},nil
}
