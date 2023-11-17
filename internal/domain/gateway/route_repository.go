package gateway

import "freigth_service/internal/domain/entity"

type RouteRepositoryInterface interface{
	Create(route *entity.Route) error
	FindByID(id string) (*entity.Route,error)
	Update(route *entity.Route) error
}