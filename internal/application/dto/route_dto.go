package dto

import (
	customtime "freigth_service/pkg/custom_time"
)

type InputCreateRouteDto struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Distance float64 `json:"distance"`
	Event    string  `json:"event"`
}

type OutputCreateRouteDto struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Distance     float64 `json:"distance"`
	Status       string  `json:"status"`
	FreightPrice float64 `json:"freight_price"`
}

type InputChangeStatusRouteDto struct {
	ID         string                `json:"id"`
	StartedAt  customtime.CustomTime `json:"started_at"`
	FinishedAt customtime.CustomTime `json:"finished_at"`
	Event      string                `json:"event"`
}

type OutputChangeStatusRouteDto struct {
	ID         string                `json:"id"`
	Status     string                `json:"status"`
	StartedAt  customtime.CustomTime `json:"started_at"`
	FinishedAt customtime.CustomTime `json:"finished_at"`
}
