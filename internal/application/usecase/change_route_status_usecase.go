package usecase

import (
	"freigth_service/internal/application/dto"
	"freigth_service/internal/domain/gateway"
	customtime "freigth_service/pkg/custom_time"
	"time"
)

type ChangeStatusRouteUseCase struct {
	RouteRepository gateway.RouteRepositoryInterface
}

func NewChangeStatusRouteUseaCase(repo gateway.RouteRepositoryInterface) *ChangeStatusRouteUseCase {
	return &ChangeStatusRouteUseCase{
		RouteRepository: repo,
	}
}

func (u *ChangeStatusRouteUseCase) Execute(input dto.InputChangeStatusRouteDto) (*dto.OutputChangeStatusRouteDto, error) {
	route, err := u.RouteRepository.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	if input.Event == "RouteStarted" {
		route.Start(time.Time(input.StartedAt))
	}
	if input.Event == "RouteFinished" {
		route.Finish(time.Time(input.FinishedAt))
	}

	err = u.RouteRepository.Update(route)
	if err != nil {
		return nil, err
	}

	return &dto.OutputChangeStatusRouteDto{
		ID:         route.ID,
		Status:     route.Status,
		StartedAt:  customtime.CustomTime(route.StartedAt),
		FinishedAt: customtime.CustomTime(route.FinishedAt),
	}, nil
}
