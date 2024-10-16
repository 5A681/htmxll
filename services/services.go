package services

import (
	"htmxll/dto"
	"htmxll/entity"
	"htmxll/filter"
	"htmxll/repository"
)

type Service interface {
	GetLatestData(bayId int, ttime string) ([]dto.DataTmps, error)
	GetDataLatestMonthDayTime(ttime string, bayId int, filter filter.SortData) ([]dto.DataTmps, error)
	GetDataLatestMonthNightTime(ttime string, bayId int, filter filter.SortData) ([]dto.DataTmps, error)
	GetDataLatestMonthAllTime(ttime string, bayId int, filter filter.SortData) ([]dto.DataTmps, error)
	GetDataLatestYearPeakTime(ttime string, bayId int, year int, filter filter.SortData) ([]dto.DataTmpsYear, error)
	GetDataLatestYearLightTime(ttime string, bayId int, year int, filter filter.SortData) ([]dto.DataTmpsYear, error)
	GetAllBay() ([]entity.Bay, error)
	GetAllBayByStationId(stationId int) ([]entity.Bay, error)
	GetAllSubStation() ([]entity.SubStation, error)
	GetFirstSubstation() (*entity.SubStation, error)
	GetBayById(bayId int) (*entity.Bay, error)
	GetSubStationById(sId int) (*entity.SubStation, error)
	GetLatestYear() (int, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return service{repo}
}

func (s service) GetLatestYear() (int, error) {
	year, err := s.repo.GetLatestYear()
	if err != nil {
		return 0, err
	}
	return year, nil
}
