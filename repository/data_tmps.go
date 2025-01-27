package repository

import (
	"fmt"
	"htmxll/entity"
	"htmxll/filter"
	"log"
	"time"
)

func (s repository) GetDataTmpsById(id int) (*entity.DataTmps, error) {
	var dataTmps entity.DataTmps
	err := s.db.Get(&dataTmps, `select * from data_tmps where id = $1 order by id asc`, id)
	if err != nil {
		return nil, err
	}
	return &dataTmps, nil
}
func (s repository) CreateDataTmep(data entity.DataTmps) error {
	sqlCreate := `insert into data_tmps (current_phase_a,current_phase_b,current_phase_c,
	active_power,reactive_power,power_factor,data_datetime,created_at,bay_id,voltage_bc,voltage_ab,voltage_ca) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
	_, err := s.db.Exec(sqlCreate, data.CurrentPhaseA, data.CurrentPhaseB, data.CurrentPhaseC, data.ActivePower, data.ReactivePower,
		data.PowerFactor, data.DataDatetime, data.CreatedAt, data.BayId, data.VoltageBC, data.VoltageBC, data.VoltageCA)
	if err != nil {
		return err
	}
	return nil
}
func (s repository) GetDataTempsByBayId(bayId int, sort filter.SortData) ([]entity.DataTmps, error) {
	var dataTemps []entity.DataTmps

	if sort.Time {
		err := s.db.Select(&dataTemps, `select * from data_tmps where bay_id = $1 order by data_datetime desc`, bayId)
		if err != nil {
			return nil, err
		}
	} else if !sort.Time {
		err := s.db.Select(&dataTemps, `select * from data_tmps where bay_id = $1 order by data_datetime asc`, bayId)
		if err != nil {
			return nil, err
		}
	}

	return dataTemps, nil
}

func (s repository) GetLatestDataByBayId(bayId int, date time.Time) ([]entity.DataTmps, error) {
	var dataTemps []entity.DataTmps

	query := fmt.Sprintf(`select * from data_tmps where bay_id = %d and date(data_datetime) = date('%s') order by data_datetime asc`, bayId, date.Format("2006-01-02"))
	log.Println(query)

	err := s.db.Select(&dataTemps, query)
	if err != nil {
		return nil, err
	}

	return dataTemps, nil

}

func (s repository) GetMaxDataPerDayPerTime(bayId int, minTime time.Time, maxTime time.Time) (*entity.DataTmps, error) {
	var dataTemps entity.DataTmps
	query := fmt.Sprintf(`select * from data_tmps where bay_id = %d and data_datetime  between '%s' and '%s'
		order by active_power desc,data_datetime asc `, bayId, minTime.Format(`2006-01-02 15:04:05`), maxTime.Format(`2006-01-02 15:04:05`))

	err := s.db.Get(&dataTemps, query)
	if err != nil {
		return nil, err
	}

	return &dataTemps, nil
}

func (s repository) GetMinDataPerDayPerTime(bayId int, minTime time.Time, maxTime time.Time) (*entity.DataTmps, error) {
	var dataTemps entity.DataTmps
	query := fmt.Sprintf(`select * from data_tmps where bay_id = %d and data_datetime  between '%s' and '%s'
		order by active_power asc,data_datetime asc `, bayId, minTime.Format(`2006-01-02 15:04:05`), maxTime.Format(`2006-01-02 15:04:05`))

	err := s.db.Get(&dataTemps, query)
	if err != nil {
		return nil, err
	}

	return &dataTemps, nil
}
func (s repository) GetMaxDataPerMonth(bayId int, year int, month int) (*entity.DataTmps, error) {
	var dataTemps entity.DataTmps
	query := `select * from  data_tmps dt where dt.bay_id= $1 and EXTRACT(YEAR FROM dt.data_datetime) = $2 and EXTRACT(MONTH FROM dt.data_datetime) = $3 order  by active_power desc,data_datetime asc limit 1`
	err := s.db.Get(&dataTemps, query, bayId, year, month)
	if err != nil {
		return nil, err
	}

	return &dataTemps, nil
}

func (s repository) GetMinDataPerMonth(bayId int, year int, month int) (*entity.DataTmps, error) {
	var dataTemps entity.DataTmps
	query := fmt.Sprintf(`select * from  data_tmps dt where dt.bay_id= %d and EXTRACT(YEAR FROM dt.data_datetime) = %d and EXTRACT(MONTH FROM dt.data_datetime) = %d order by active_power asc,data_datetime asc limit 1`, bayId, year, month)

	err := s.db.Get(&dataTemps, query)
	if err != nil {
		return nil, err
	}

	return &dataTemps, nil
}

func (s repository) GetMaxDataPerDayPerTimeTwoTime(bayId int, minTime1 time.Time, maxTime1 time.Time, minTime2 time.Time, maxTime2 time.Time) (*entity.DataTmps, error) {
	var dataTemps entity.DataTmps
	// query := `select * from data_tmps where bay_id = $1 and data_datetime  between $2 and $3 or data_datetime  between $4 and $5
	// 	order by active_power desc,data_datetime asc`
	//query := "select * from data_tmps where bay_id = $1"
	err := s.db.Get(&dataTemps, `select * from data_tmps where bay_id = $1 and (data_datetime between $2 and $3 or data_datetime between $4 and $5 ) 
		order by active_power desc,data_datetime asc`, bayId, minTime1, maxTime1, minTime2, maxTime2)
	if err != nil {
		return nil, err
	}

	return &dataTemps, nil
}
func (s repository) GetMinDataPerDayPerTimeTwoTime(bayId int, minTime1 time.Time, maxTime1 time.Time, minTime2 time.Time, maxTime2 time.Time) (*entity.DataTmps, error) {
	var dataTemps entity.DataTmps
	// query := `select * from data_tmps where bay_id = $1 and data_datetime  between $2 and $3 or data_datetime  between $4 and $5
	// 	order by active_power desc,data_datetime asc`
	//query := "select * from data_tmps where bay_id = $1"
	err := s.db.Get(&dataTemps, `select * from data_tmps where bay_id = $1 and (data_datetime between $2 and $3 or data_datetime between $4 and $5 ) 
		order by active_power asc,data_datetime asc`, bayId, minTime1, maxTime1, minTime2, maxTime2)
	if err != nil {
		return nil, err
	}

	return &dataTemps, nil
}

func (s repository) GetLatestYear() (int, error) {
	type Result struct {
		Year int `db:"year"`
	}
	r := Result{}
	err := s.db.Get(&r, `select max(extract(year from data_datetime)) as year from data_tmps dt`)
	if err != nil {
		return 0, err
	}
	return r.Year, nil
}
