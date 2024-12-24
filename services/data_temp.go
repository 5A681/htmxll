package services

import (
	"fmt"
	"htmxll/config"
	"htmxll/dto"
	"htmxll/entity"
	"htmxll/filter"
	"log"
	"time"
)

func (s service) GetLatestData(bayId int, ttime time.Time) ([]dto.DataTmps, error) {

	if ttime.IsZero() {
		maxdate, err := s.repo.GetMaxDate()
		if err != nil {
			return nil, err
		}

		if maxdate != nil {
			datas, err := s.repo.GetLatestDataByBayId(bayId, *maxdate)
			if err != nil {
				return nil, err
			}

			var res []dto.DataTmps
			for _, data := range datas {
				res = append(res, dto.DataTmps{
					Id:            data.Id,
					CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
					CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
					CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
					ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
					ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
					PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
					Date:          data.DataDatetime.Format("02/01/2006"),
					Time:          data.DataDatetime.Format("15:04"),
					BayId:         data.BayId,
					CreatedAt:     data.CreatedAt,
					Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
					Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
					Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
				})
			}
			return res, nil
		}
	}

	maxdate := ttime
	datas, err := s.repo.GetLatestDataByBayId(bayId, maxdate)
	if err != nil {
		return nil, err
	}

	var res []dto.DataTmps
	for _, data := range datas {
		res = append(res, dto.DataTmps{
			Id:            data.Id,
			CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
			CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
			CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
			ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
			ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
			PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
			Date:          data.DataDatetime.Format("02/01/2006"),
			Time:          data.DataDatetime.Format("15:04"),
			BayId:         data.BayId,
			CreatedAt:     data.CreatedAt,
			Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
			Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
			Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
		})
	}
	return res, nil

}

func (s service) GetDataLatestMonthDayTime(ttime time.Time, bayId int, filter filter.SortData) ([]dto.DataTmps, error) {
	var res []dto.DataTmps
	err := s.repo.CheckPreviousMonth()
	if err != nil {
		log.Println(err.Error())
		currentTime, err := s.repo.GetMaxDate()
		if err != nil {
			return nil, err
		}
		for i := 0; i < 31; i++ {

			if ttime.IsZero() {
				min := time.Date(currentTime.Year(), currentTime.Month(), i+1, 8, 0, 0, 0, time.Local)
				max := time.Date(currentTime.Year(), currentTime.Month(), i+1, 15, 30, 0, 0, time.Local)
				data, err := s.repo.GetMaxDataPerDayPerTime(bayId, min, max)
				if err != nil {
					log.Println(err.Error())
				}
				if data != nil {
					res = append(res, dto.DataTmps{
						Id:            data.Id,
						CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
						CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
						CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
						ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
						ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
						PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
						Date:          data.DataDatetime.Format("02/01/2006"),
						Time:          data.DataDatetime.Format("15:04"),
						BayId:         data.BayId,
						CreatedAt:     data.CreatedAt,
						Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
						Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
						Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
					})

				}
			} else {

				maxdate := ttime

				min := time.Date(maxdate.Year(), maxdate.Month(), i+1, 8, 0, 0, 0, time.Local)
				max := time.Date(maxdate.Year(), maxdate.Month(), i+1, 15, 30, 0, 0, time.Local)
				data, err := s.repo.GetMaxDataPerDayPerTime(bayId, min, max)
				if err != nil {
					log.Println(err.Error())
				}
				if data != nil {
					res = append(res, dto.DataTmps{
						Id:            data.Id,
						CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
						CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
						CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
						ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
						ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
						PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
						Date:          data.DataDatetime.Format("02/01/2006"),
						Time:          data.DataDatetime.Format("15:04"),
						BayId:         data.BayId,
						CreatedAt:     data.CreatedAt,
						Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
						Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
						Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
					})

				}
			}

		}

	} else {
		currentTime := time.Now()
		currentTime = currentTime.AddDate(0, -1, 0)
		for i := 0; i < 31; i++ {
			min := time.Date(currentTime.Year(), currentTime.Month(), i+1, 8, 0, 0, 0, time.Local)
			max := time.Date(currentTime.Year(), currentTime.Month(), i+1, 15, 30, 0, 0, time.Local)
			data, err := s.repo.GetMaxDataPerDayPerTime(bayId, min, max)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {
				res = append(res, dto.DataTmps{
					Id:            data.Id,
					CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
					CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
					CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
					ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
					ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
					PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
					Date:          data.DataDatetime.Format("02/01/2006"),
					Time:          data.DataDatetime.Format("15:04"),
					BayId:         data.BayId,
					CreatedAt:     data.CreatedAt,
					Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
					Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
					Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
				})
			}
		}
	}

	return res, nil

}

func (s service) GetDataLatestMonthNightTime(ttime time.Time, bayId int, filter filter.SortData) ([]dto.DataTmps, error) {
	var res []dto.DataTmps
	err := s.repo.CheckPreviousMonth()
	if err != nil {
		log.Println(err.Error())
		currentTime, err := s.repo.GetMaxDate()
		if err != nil {
			return nil, err
		}
		for i := 0; i < 31; i++ {
			if ttime.IsZero() {
				min := time.Date(currentTime.Year(), currentTime.Month(), i+1, 0, 0, 0, 0, time.Local)
				max := time.Date(currentTime.Year(), currentTime.Month(), i+1, 7, 30, 0, 0, time.Local)
				min2 := time.Date(currentTime.Year(), currentTime.Month(), i+1, 16, 0, 0, 0, time.Local)
				max2 := time.Date(currentTime.Year(), currentTime.Month(), i+1, 23, 30, 0, 0, time.Local)
				data, err := s.repo.GetMaxDataPerDayPerTimeTwoTime(bayId, min, max, min2, max2)
				if err != nil {
					log.Println(err.Error(), i)
				}
				if data != nil {
					res = append(res, dto.DataTmps{
						Id:            data.Id,
						CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
						CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
						CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
						ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
						ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
						PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
						Date:          data.DataDatetime.Format("02/01/2006"),
						Time:          data.DataDatetime.Format("15:04"),
						BayId:         data.BayId,
						CreatedAt:     data.CreatedAt,
						Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
						Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
						Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
					})
				}
			} else {

				maxdate := ttime

				min := time.Date(maxdate.Year(), maxdate.Month(), i+1, 0, 0, 0, 0, time.Local)
				max := time.Date(maxdate.Year(), maxdate.Month(), i+1, 7, 30, 0, 0, time.Local)
				min2 := time.Date(maxdate.Year(), maxdate.Month(), i+1, 16, 0, 0, 0, time.Local)
				max2 := time.Date(maxdate.Year(), maxdate.Month(), i+1, 23, 30, 0, 0, time.Local)
				data, err := s.repo.GetMaxDataPerDayPerTimeTwoTime(bayId, min, max, min2, max2)
				if err != nil {
					log.Println(err.Error(), i)
				}
				if data != nil {
					res = append(res, dto.DataTmps{
						Id:            data.Id,
						CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
						CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
						CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
						ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
						ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
						PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
						Date:          data.DataDatetime.Format("02/01/2006"),
						Time:          data.DataDatetime.Format("15:04"),
						BayId:         data.BayId,
						CreatedAt:     data.CreatedAt,
						Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
						Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
						Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
					})
				}
			}
		}

	} else {
		currentTime := time.Now()
		currentTime = currentTime.AddDate(0, -1, 0)
		for i := 0; i < 31; i++ {
			min := time.Date(currentTime.Year(), currentTime.Month(), i+1, 0, 0, 0, 0, time.Local)
			max := time.Date(currentTime.Year(), currentTime.Month(), i+1, 7, 30, 0, 0, time.Local)
			min2 := time.Date(currentTime.Year(), currentTime.Month(), i+1, 16, 0, 0, 0, time.Local)
			max2 := time.Date(currentTime.Year(), currentTime.Month(), i+1, 23, 30, 0, 0, time.Local)
			data, err := s.repo.GetMaxDataPerDayPerTimeTwoTime(bayId, min, max, min2, max2)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {
				res = append(res, dto.DataTmps{
					Id:            data.Id,
					CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
					CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
					CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
					ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
					ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
					PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
					Date:          data.DataDatetime.Format("02/01/2006"),
					Time:          data.DataDatetime.Format("15:04"),
					BayId:         data.BayId,
					CreatedAt:     data.CreatedAt,
					Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
					Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
					Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
				})
			}
		}
	}

	return res, nil

}
func (s service) GetDataLatestMonthAllTime(ttime time.Time, bayId int, filter filter.SortData) ([]dto.DataTmps, error) {
	var res []dto.DataTmps
	err := s.repo.CheckPreviousMonth()
	if err != nil {
		log.Println(err.Error())
		currentTime, err := s.repo.GetMaxDate()
		if err != nil {
			return nil, err
		}
		for i := 0; i < 31; i++ {
			if ttime.IsZero() {
				min := time.Date(currentTime.Year(), currentTime.Month(), i+1, 0, 0, 0, 0, time.Local)
				max := time.Date(currentTime.Year(), currentTime.Month(), i+1, 23, 30, 0, 0, time.Local)
				data, err := s.repo.GetMinDataPerDayPerTime(bayId, min, max)
				if err != nil {
					log.Println(err.Error())
				}
				if data != nil {
					res = append(res, dto.DataTmps{
						Id:            data.Id,
						CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
						CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
						CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
						ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
						ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
						PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
						Date:          data.DataDatetime.Format("02/01/2006"),
						Time:          data.DataDatetime.Format("15:04"),
						BayId:         data.BayId,
						CreatedAt:     data.CreatedAt,
						Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
						Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
						Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
					})
				}
			} else {
				maxdate := ttime
				min := time.Date(maxdate.Year(), maxdate.Month(), i+1, 0, 0, 0, 0, time.Local)
				max := time.Date(maxdate.Year(), maxdate.Month(), i+1, 23, 30, 0, 0, time.Local)
				data, err := s.repo.GetMinDataPerDayPerTime(bayId, min, max)
				if err != nil {
					log.Println(err.Error())
				}
				if data != nil {
					res = append(res, dto.DataTmps{
						Id:            data.Id,
						CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
						CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
						CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
						ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
						ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
						PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
						Date:          data.DataDatetime.Format("02/01/2006"),
						Time:          data.DataDatetime.Format("15:04"),
						BayId:         data.BayId,
						CreatedAt:     data.CreatedAt,
						Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
						Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
						Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
					})
				}
			}
		}

	} else {
		currentTime := time.Now()
		currentTime = currentTime.AddDate(0, -1, 0)
		for i := 0; i < 31; i++ {
			min := time.Date(currentTime.Year(), currentTime.Month(), i+1, 0, 0, 0, 0, time.Local)
			max := time.Date(currentTime.Year(), currentTime.Month(), i+1, 23, 30, 0, 0, time.Local)
			data, err := s.repo.GetMinDataPerDayPerTime(bayId, min, max)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {
				res = append(res, dto.DataTmps{
					Id:            data.Id,
					CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
					CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
					CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
					ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
					ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
					PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
					Date:          data.DataDatetime.Format("02/01/2006"),
					Time:          data.DataDatetime.Format("15:04"),
					BayId:         data.BayId,
					CreatedAt:     data.CreatedAt,
					Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
					Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
					Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
				})
			}
		}
	}

	return res, nil

}

func (s service) GetDataLatestYearPeakTime(ttime *time.Time, bayId int, year int, filter filter.SortData) ([]dto.DataTmpsYear, error) {
	var datas []dto.DataTmpsYear
	for i := 0; i < 12; i++ {

		if ttime.IsZero() {
			data, err := s.repo.GetMaxDataPerMonth(bayId, year, i+1)
			if err != nil {
				log.Println(err.Error(), 1)
			}
			if data != nil {
				ttime = &data.DataDatetime
				datas = append(datas, dto.DataTmpsYear{
					Id:            data.Id,
					Month:         data.DataDatetime.Format("Jan"),
					Date:          data.DataDatetime.Format("02"),
					Time:          data.DataDatetime.Format("15:04"),
					CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
					CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
					CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
					ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
					ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
					PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
					CreatedAt:     data.CreatedAt,
					BayId:         data.BayId,
					Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
					Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
					Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
				})
			}
		} else {
			maxdate := ttime
			data, err := s.repo.GetMaxDataPerMonth(bayId, maxdate.Year(), i+1)
			if err != nil {
				log.Println(err.Error(), 1)
			}
			if data != nil {
				datas = append(datas, dto.DataTmpsYear{
					Id:            data.Id,
					Month:         data.DataDatetime.Format("Jan"),
					Date:          data.DataDatetime.Format("02"),
					Time:          data.DataDatetime.Format("15:04"),
					CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
					CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
					CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
					ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
					ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
					PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
					CreatedAt:     data.CreatedAt,
					BayId:         data.BayId,
					Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
					Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
					Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
				})
			}
		}
	}

	return datas, nil

}

func (s service) GetDataLatestYearLightTime(ttime *time.Time, bayId int, year int, filter filter.SortData) ([]dto.DataTmpsYear, error) {

	var datas []dto.DataTmpsYear

	for i := 0; i < 12; i++ {

		if ttime.IsZero() {
			data, err := s.repo.GetMinDataPerMonth(bayId, year, i+1)
			if err != nil {
				log.Println(err.Error(), 3)
			}
			if data != nil {
				ttime = &data.DataDatetime
				datas = append(datas, dto.DataTmpsYear{
					Id:            data.Id,
					Month:         data.DataDatetime.Format("Jan"),
					Date:          data.DataDatetime.Format("02"),
					Time:          data.DataDatetime.Format("15:04"),
					CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
					CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
					CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
					ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
					ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
					PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
					CreatedAt:     data.CreatedAt,
					BayId:         data.BayId,
					Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
					Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
					Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
				})
			}
		} else {
			maxdate := ttime
			data, err := s.repo.GetMinDataPerMonth(bayId, maxdate.Year(), i+1)
			if err != nil {
				log.Println(err.Error(), 3)
			}
			if data != nil {
				datas = append(datas, dto.DataTmpsYear{
					Id:            data.Id,
					Month:         data.DataDatetime.Format("Jan"),
					Date:          data.DataDatetime.Format("02"),
					Time:          data.DataDatetime.Format("15:04"),
					CurrentPhaseA: fmt.Sprintf("%.2f", data.CurrentPhaseA),
					CurrentPhaseB: fmt.Sprintf("%.2f", data.CurrentPhaseB),
					CurrentPhaseC: fmt.Sprintf("%.2f", data.CurrentPhaseC),
					ActivePower:   fmt.Sprintf("%.2f", data.ActivePower),
					ReactivePower: fmt.Sprintf("%.2f", data.ReactivePower),
					PowerFactor:   fmt.Sprintf("%.2f", data.PowerFactor),
					CreatedAt:     data.CreatedAt,
					BayId:         data.BayId,
					Vab:           fmt.Sprintf("%.2f", data.VoltageAB),
					Vbc:           fmt.Sprintf("%.2f", data.VoltageBC),
					Vca:           fmt.Sprintf("%.2f", data.VoltageCA),
				})
			}
		}
	}

	return datas, nil

}

func (s service) GetAllBayByStationId(config config.Config, stationId int) ([]entity.Bay, error) {
	res, err := s.repo.GetBaysByStationId(stationId)
	if err != nil {
		return nil, err
	}
	bays := []entity.Bay{}
	for _, bay := range res {
		if bay.Name == "TP1" || bay.Name == "TP2" {
			bay.Name = "line " + config.LINE_KV + "/" + bay.Name
		}
		bays = append(bays, entity.Bay{
			Id:           bay.Id,
			Name:         bay.Name,
			CreatedAt:    bay.CreatedAt,
			SubStationId: bay.SubStationId,
		})
	}

	return bays, nil
}

func (s service) GetAllBay() ([]entity.Bay, error) {
	res, err := s.repo.GetBays()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) GetAllSubStation() ([]entity.SubStation, error) {
	res, err := s.repo.GetSubStations()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) GetFirstSubstation() (*entity.SubStation, error) {
	res, err := s.repo.GetFirstSubstation()
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s service) GetBayById(bayId int) (*entity.Bay, error) {
	res, err := s.repo.GetBayById(bayId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s service) GetSubStationById(sId int) (*entity.SubStation, error) {
	res, err := s.repo.GetSubStationById(sId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
