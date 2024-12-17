package services

import (
	"htmxll/config"
	"htmxll/dto"
	"htmxll/entity"
	"htmxll/filter"
	"htmxll/repository"
	"time"
)

type Service interface {
	GetLatestData(bayId int, ttime time.Time) ([]dto.DataTmps, error)
	GetDataLatestMonthDayTime(ttime time.Time, bayId int, filter filter.SortData) ([]dto.DataTmps, error)
	GetDataLatestMonthNightTime(ttime time.Time, bayId int, filter filter.SortData) ([]dto.DataTmps, error)
	GetDataLatestMonthAllTime(ttime time.Time, bayId int, filter filter.SortData) ([]dto.DataTmps, error)
	GetDataLatestYearPeakTime(ttime *time.Time, bayId int, year int, filter filter.SortData) ([]dto.DataTmpsYear, error)
	GetDataLatestYearLightTime(ttime *time.Time, bayId int, year int, filter filter.SortData) ([]dto.DataTmpsYear, error)
	GetAllBay() ([]entity.Bay, error)
	GetAllBayByStationId(config config.Config, stationId int) ([]entity.Bay, error)
	GetAllSubStation() ([]entity.SubStation, error)
	GetFirstSubstation() (*entity.SubStation, error)
	GetBayById(bayId int) (*entity.Bay, error)
	GetSubStationById(sId int) (*entity.SubStation, error)
	GetLatestYear() (int, error)
	GetRowsMonthlyData(config config.Config, bayId int, ttime *time.Time) ([]dto.MonthlyRowData, error)
}

type service struct {
	repo   repository.Repository
	config config.Config
}

func NewService(repo repository.Repository, config config.Config) Service {
	return service{repo, config}
}

func (s service) GetLatestYear() (int, error) {
	year, err := s.repo.GetLatestYear()
	if err != nil {
		return 0, err
	}
	return year, nil
}
