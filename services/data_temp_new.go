package services

import (
	"fmt"
	"htmxll/config"
	"htmxll/dto"
	"log"
	"time"
)

func (s service) GetRowsMonthlyData(config config.Config, bayId int, ttime *time.Time) ([]dto.MonthlyRowData, error) {
	var res []dto.MonthlyRowData
	if bayId == 0 {
		return res, nil
	}
	if ttime != nil {
		for i := 1; i <= getMaxDayMonth(ttime.Year(), ttime.Month()); i++ {
			day, night, all, err := s.GetDataTable(config, bayId, ttime, i)
			if err != nil {
				log.Println(err)
			}
			_ = day
			_ = night
			_ = all
			res = append(res, dto.MonthlyRowData{
				PeakDay:   day,
				PeakNight: night,
				All:       all,
			})
		}

	}
	return res, nil
}

func (s service) GetDataTable(config config.Config, bayId int, ttime *time.Time, i int) (dto.MonthlyData, dto.MonthlyData, dto.MonthlyData, error) {
	var peakDay dto.MonthlyData
	var lightNight dto.MonthlyData
	var all dto.MonthlyData
	dayTimeMin := time.Date(ttime.Year(), ttime.Month(), i, 8, 0, 0, 0, time.Local)
	dayTimeMax := time.Date(ttime.Year(), ttime.Month(), i, 15, 30, 0, 0, time.Local)
	day, _ := s.repo.GetMaxDataPerDayPerTime(bayId, dayTimeMin, dayTimeMax)

	if day != nil {
		peakDay = dto.MonthlyData{
			Date: day.DataDatetime.Format("02/01/2006"),
			Time: day.DataDatetime.Format("15:04"),
			Vab:  fmt.Sprintf("%.2f", day.VoltageAB),
			Vbc:  fmt.Sprintf("%.2f", day.VoltageBC),
			Vca:  fmt.Sprintf("%.2f", day.VoltageCA),
			Ia:   fmt.Sprintf("%.2f", day.CurrentPhaseA),
			Ib:   fmt.Sprintf("%.2f", day.CurrentPhaseB),
			Ic:   fmt.Sprintf("%.2f", day.CurrentPhaseC),
			Mw:   fmt.Sprintf("%.2f", day.ActivePower),
			Mvar: fmt.Sprintf("%.2f", day.ReactivePower),
			P:    fmt.Sprintf("%.2f", day.PowerFactor),
		}
	} else {
		timeDay := time.Date(ttime.Year(), ttime.Month(), i, 0, 0, 0, 0, time.Local)
		peakDay = dto.MonthlyData{
			Date: timeDay.Format("02/01/2006"),
			Time: timeDay.Format("15:04"),
			Vab:  "0.00",
			Vbc:  "0.00",
			Vca:  "0.00",
			Ia:   "0.00",
			Ib:   "0.00",
			Ic:   "0.00",
			Mw:   "0.00",
			Mvar: "0.00",
			P:    "0.00",
		}
	}
	dayTimeMin1 := time.Date(ttime.Year(), ttime.Month(), i, 0, 0, 0, 0, time.Local)
	dayTimeMax1 := time.Date(ttime.Year(), ttime.Month(), i, 7, 30, 0, 0, time.Local)
	dayTimeMin2 := time.Date(ttime.Year(), ttime.Month(), i, 16, 30, 0, 0, time.Local)
	dayTimeMax2 := time.Date(ttime.Year(), ttime.Month(), i, 23, 30, 0, 0, time.Local)
	night, _ := s.repo.GetMaxDataPerDayPerTimeTwoTime(bayId, dayTimeMin1, dayTimeMax1, dayTimeMin2, dayTimeMax2)

	if night != nil {
		lightNight = dto.MonthlyData{
			Date: night.DataDatetime.Format("02/01/2006"),
			Time: night.DataDatetime.Format("15:04"),
			Vab:  fmt.Sprintf("%.2f", night.VoltageAB),
			Vbc:  fmt.Sprintf("%.2f", night.VoltageBC),
			Vca:  fmt.Sprintf("%.2f", night.VoltageCA),
			Ia:   fmt.Sprintf("%.2f", night.CurrentPhaseA),
			Ib:   fmt.Sprintf("%.2f", night.CurrentPhaseB),
			Ic:   fmt.Sprintf("%.2f", night.CurrentPhaseC),
			Mw:   fmt.Sprintf("%.2f", night.ActivePower),
			Mvar: fmt.Sprintf("%.2f", night.ReactivePower),
			P:    fmt.Sprintf("%.2f", night.PowerFactor),
		}
	} else {
		timeDay := time.Date(ttime.Year(), ttime.Month(), i, 0, 0, 0, 0, time.Local)
		lightNight = dto.MonthlyData{
			Date: timeDay.Format("02/01/2006"),
			Time: timeDay.Format("15:04"),
			Vab:  "0.00",
			Vbc:  "0.00",
			Vca:  "0.00",
			Ia:   "0.00",
			Ib:   "0.00",
			Ic:   "0.00",
			Mw:   "0.00",
			Mvar: "0.00",
			P:    "0.00",
		}
	}
	dayTimeMin = time.Date(ttime.Year(), ttime.Month(), i, 0, 0, 0, 0, time.Local)
	dayTimeMax = time.Date(ttime.Year(), ttime.Month(), i, 23, 30, 0, 0, time.Local)
	allData, _ := s.repo.GetMinDataPerDayPerTime(bayId, dayTimeMin, dayTimeMax)

	if allData != nil {
		all = dto.MonthlyData{
			Date: allData.DataDatetime.Format("02/01/2006"),
			Time: allData.DataDatetime.Format("15:04"),
			Vab:  fmt.Sprintf("%.2f", allData.VoltageAB),
			Vbc:  fmt.Sprintf("%.2f", allData.VoltageBC),
			Vca:  fmt.Sprintf("%.2f", allData.VoltageCA),
			Ia:   fmt.Sprintf("%.2f", allData.CurrentPhaseA),
			Ib:   fmt.Sprintf("%.2f", allData.CurrentPhaseB),
			Ic:   fmt.Sprintf("%.2f", allData.CurrentPhaseC),
			Mw:   fmt.Sprintf("%.2f", allData.ActivePower),
			Mvar: fmt.Sprintf("%.2f", allData.ReactivePower),
			P:    fmt.Sprintf("%.2f", allData.PowerFactor),
		}
	} else {
		timeDay := time.Date(ttime.Year(), ttime.Month(), i, 0, 0, 0, 0, time.Local)
		all = dto.MonthlyData{
			Date: timeDay.Format("02/01/2006"),
			Time: timeDay.Format("15:04"),
			Vab:  "0.00",
			Vbc:  "0.00",
			Vca:  "0.00",
			Ia:   "0.00",
			Ib:   "0.00",
			Ic:   "0.00",
			Mw:   "0.00",
			Mvar: "0.00",
			P:    "0.00",
		}
	}
	return peakDay, lightNight, all, nil
}

func getMaxDayMonth(year int, month time.Month) int {
	// The zero value of day 0 in the next month gives the last day of the current month
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
