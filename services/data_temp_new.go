package services

import (
	"htmxll/dto"
	"log"
	"time"
)

func (s service) GetRowsMonthlyData(ttime time.Time) ([]dto.MonthlyRowData, error) {
	bays, err := s.GetAllBay()
	if err != nil {
		return nil, err
	}
	res := []dto.MonthlyRowData{}
	for _, item := range bays {
		day, err := s.GetNewMonthlyPeakDay(item.Id, ttime)
		if err != nil {
			log.Println("error :", err)
			return nil, err
		}
		night, err := s.GetNewMonthlyLowNight(item.Id, ttime)
		if err != nil {
			log.Println("error :", err)
			return nil, err
		}
		all, err := s.GetNewMonthlyLowAll(item.Id, ttime)
		if err != nil {
			log.Println("error :", err)
			return nil, err
		}
		res = append(res, dto.MonthlyRowData{
			Bay:       item.Name,
			PeakDay:   *day,
			PeakNight: *night,
			AllLow:    *all,
		})
	}
	return res, nil
}

func (s service) GetNewMonthlyPeakDay(bayId int, ttime time.Time) (*dto.MonthlyData, error) {
	res := dto.MonthlyData{}
	err := s.repo.CheckPreviousMonth()
	if err != nil {
		log.Println(err.Error())
		currentTime, err := s.repo.GetMaxDate()
		if err != nil {
			return nil, err
		}

		if ttime.IsZero() {

			data, err := s.repo.GetMaxDataByBayIdAndMonth(bayId, currentTime.Year(), int(currentTime.Month()), 8, 15)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {
				res.Date = data.DataDatetime.Format("02/01/2006")
				res.Time = data.DataDatetime.Format("15:04")
				res.Kv = data.VoltageBC
				res.Ia = data.CurrentPhaseA
				res.Ib = data.CurrentPhaseB
				res.Ic = data.CurrentPhaseC
				res.Mw = data.ActivePower
				res.Mvar = data.ReactivePower

			}
		} else {

			maxdate := ttime

			data, err := s.repo.GetMaxDataByBayIdAndMonth(bayId, maxdate.Year(), int(maxdate.Month()), 8, 15)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {
				res.Date = data.DataDatetime.Format("02/01/2006")
				res.Time = data.DataDatetime.Format("15:04")
				res.Kv = data.VoltageBC
				res.Ia = data.CurrentPhaseA
				res.Ib = data.CurrentPhaseB
				res.Ic = data.CurrentPhaseC
				res.Mw = data.ActivePower
				res.Mvar = data.ReactivePower

			}
		}

	} else {
		currentTime := time.Now()
		currentTime = currentTime.AddDate(0, -1, 0)

		data, err := s.repo.GetMaxDataByBayIdAndMonth(bayId, currentTime.Year(), int(currentTime.Month()), 8, 15)
		if err != nil {
			log.Println(err.Error())
		}
		if data != nil {
			res.Date = data.DataDatetime.Format("02/01/2006")
			res.Time = data.DataDatetime.Format("15:04")
			res.Kv = data.VoltageBC
			res.Ia = data.CurrentPhaseA
			res.Ib = data.CurrentPhaseB
			res.Ic = data.CurrentPhaseC
			res.Mw = data.ActivePower
			res.Mvar = data.ReactivePower

		}
	}
	return &res, nil
}

func (s service) GetNewMonthlyLowNight(bayId int, ttime time.Time) (*dto.MonthlyData, error) {
	res := dto.MonthlyData{}
	err := s.repo.CheckPreviousMonth()
	if err != nil {
		log.Println(err.Error())
		currentTime, err := s.repo.GetMaxDate()
		if err != nil {
			return nil, err
		}

		if ttime.IsZero() {

			data, err := s.repo.GetMinDataByBayIdAndMonth(bayId, currentTime.Year(), int(currentTime.Month()), 16, 23)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {
				res.Date = data.DataDatetime.Format("02/01/2006")
				res.Time = data.DataDatetime.Format("15:04")
				res.Kv = data.VoltageBC
				res.Ia = data.CurrentPhaseA
				res.Ib = data.CurrentPhaseB
				res.Ic = data.CurrentPhaseC
				res.Mw = data.ActivePower
				res.Mvar = data.ReactivePower

			}
		} else {

			maxdate := ttime

			data, err := s.repo.GetMinDataByBayIdAndMonth(bayId, maxdate.Year(), int(maxdate.Month()), 16, 23)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {
				res.Date = data.DataDatetime.Format("02/01/2006")
				res.Time = data.DataDatetime.Format("15:04")
				res.Kv = data.VoltageBC
				res.Ia = data.CurrentPhaseA
				res.Ib = data.CurrentPhaseB
				res.Ic = data.CurrentPhaseC
				res.Mw = data.ActivePower
				res.Mvar = data.ReactivePower

			}
		}

	} else {
		currentTime := time.Now()
		currentTime = currentTime.AddDate(0, -1, 0)

		data, err := s.repo.GetMinDataByBayIdAndMonth(bayId, currentTime.Year(), int(currentTime.Month()), 16, 23)
		if err != nil {
			log.Println(err.Error())
		}
		if data != nil {
			res.Date = data.DataDatetime.Format("02/01/2006")
			res.Time = data.DataDatetime.Format("15:04")
			res.Kv = data.VoltageBC
			res.Ia = data.CurrentPhaseA
			res.Ib = data.CurrentPhaseB
			res.Ic = data.CurrentPhaseC
			res.Mw = data.ActivePower
			res.Mvar = data.ReactivePower

		}
	}
	return &res, nil
}

func (s service) GetNewMonthlyLowAll(bayId int, ttime time.Time) (*dto.MonthlyData, error) {
	res := dto.MonthlyData{}
	err := s.repo.CheckPreviousMonth()
	if err != nil {
		log.Println(err.Error())
		currentTime, err := s.repo.GetMaxDate()
		if err != nil {
			return nil, err
		}

		if ttime.IsZero() {

			data, err := s.repo.GetMinDataByBayIdAndMonth(bayId, currentTime.Year(), int(currentTime.Month()), 0, 23)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {

				res.Date = data.DataDatetime.Format("02/01/2006")
				res.Time = data.DataDatetime.Format("15:04")
				res.Kv = data.VoltageBC
				res.Ia = data.CurrentPhaseA
				res.Ib = data.CurrentPhaseB
				res.Ic = data.CurrentPhaseC
				res.Mw = data.ActivePower
				res.Mvar = data.ReactivePower

			}
		} else {
			maxdate := ttime

			data, err := s.repo.GetMinDataByBayIdAndMonth(bayId, maxdate.Year(), int(maxdate.Month()), 0, 23)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {
				res.Date = data.DataDatetime.Format("02/01/2006")
				res.Time = data.DataDatetime.Format("15:04")
				res.Kv = data.VoltageBC
				res.Ia = data.CurrentPhaseA
				res.Ib = data.CurrentPhaseB
				res.Ic = data.CurrentPhaseC
				res.Mw = data.ActivePower
				res.Mvar = data.ReactivePower

			}
		}

	} else {
		currentTime := time.Now()
		currentTime = currentTime.AddDate(0, -1, 0)

		data, err := s.repo.GetMinDataByBayIdAndMonth(bayId, currentTime.Year(), int(currentTime.Month()), 0, 23)
		if err != nil {
			log.Println(err.Error())
		}
		if data != nil {
			res.Date = data.DataDatetime.Format("02/01/2006")
			res.Time = data.DataDatetime.Format("15:04")
			res.Kv = data.VoltageBC
			res.Ia = data.CurrentPhaseA
			res.Ib = data.CurrentPhaseB
			res.Ic = data.CurrentPhaseC
			res.Mw = data.ActivePower
			res.Mvar = data.ReactivePower

		}
	}
	return &res, nil
}
