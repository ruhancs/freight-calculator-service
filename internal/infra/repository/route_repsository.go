package repository

import (
	"database/sql"
	"freigth_service/internal/domain/entity"
)

type RouteRepository struct {
	DB *sql.DB
}

func NewRouteRepository(db *sql.DB) *RouteRepository{
	return &RouteRepository{
		DB: db,
	}
}

func(repo *RouteRepository) Create(route *entity.Route) error{
	query := "insert into routes (id,name,distance,status,freight_price) values(?,?,?,?,?)"
	_,err :=repo.DB.Exec(query,route.ID,route.Name,route.Distance,route.Status,route.FreightPrice,route.StartedAt,route.FinishedAt)
	if err != nil {
		return err
	}
	return nil
}

func(repo *RouteRepository) FindByID(id string) (*entity.Route,error){
	query := "select * from routes where id=?"
	row := repo.DB.QueryRow(query,id)
	var startedAt, finishedAt sql.NullTime
	var route entity.Route
	err := row.Scan(
		&route.ID,
		&route.Name,
		&route.Distance,
		&route.Status,
		&route.FreightPrice,
		&startedAt,
		&finishedAt,
	)
	if err != nil {
		return nil,err
	}
	if startedAt.Valid {
		route.StartedAt = startedAt.Time
	}
	if finishedAt.Valid {
		route.FinishedAt = finishedAt.Time
	}
	return &route,nil
}

func(repo *RouteRepository) Update(route *entity.Route) error{
	startedAt := route.StartedAt.Format("2006-01-02 15:04:05")
	finishedAt := route.StartedAt.Format("2006-01-02 15:04:05")
	query := "update routes set status=?, freight_price=?, started_at=?, finished_at=? where id=?"
	_,err := repo.DB.Exec(query,route.Status,route.FreightPrice,startedAt,finishedAt,route.ID)
	if err != nil {
		return err
	}
	return nil
}