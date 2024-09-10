package services

import (
	"htmxll/dto"
	"htmxll/entity"
	"htmxll/filter"
	"htmxll/repository"
)

type Service interface {
	GetLatestData(bayId int, filter filter.SortData) ([]dto.DataTmps, error)
	GetDataLatestMonthDayTime(bayId int, filter filter.SortData) ([]entity.DataTmps, error)
	GetDataLatestMonthNightTime(bayId int, filter filter.SortData) ([]entity.DataTmps, error)
	GetDataLatestMonthAllTime(bayId int, filter filter.SortData) ([]entity.DataTmps, error)
	GetDataLatestYearPeakTime(bayId int, year int, filter filter.SortData) ([]entity.DataTmps, error)
	GetDataLatestYearLightTime(bayId int, year int, filter filter.SortData) ([]entity.DataTmps, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return service{repo}
}
