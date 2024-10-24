package repository

import (
	"fmt"
	"htmxll/entity"
	"log"
)

func (s repository) GetMaxDataByBayIdAndMonth(bayId int, year int, month int, minHour int, maxHour int) (*entity.DataTmps, error) {
	var dataTemps entity.DataTmps
	query := fmt.Sprintf(`select * from data_tmps dt where bay_id = %d and extract (hour from dt.data_datetime) between %d and %d and extract( month from data_datetime) = %d and extract( year from data_datetime) = %d
		order by active_power asc,data_datetime asc `, bayId, minHour, maxHour, month, year)

	log.Println(query)

	err := s.db.Get(&dataTemps, query)
	if err != nil {
		log.Println("this error", err)
		return nil, err
	}

	return &dataTemps, nil
}

func (s repository) GetMinDataByBayIdAndMonth(bayId int, year int, month int, minHour int, maxHour int) (*entity.DataTmps, error) {
	var dataTemps entity.DataTmps
	query := fmt.Sprintf(`select * from data_tmps dt where bay_id = %d and extract (hour from dt.data_datetime) between %d and %d and extract( month from data_datetime) = %d and extract( year from data_datetime) = %d
		order by active_power asc,data_datetime asc `, bayId, minHour, maxHour, month, year)

	err := s.db.Get(&dataTemps, query)
	if err != nil {
		return nil, err
	}

	return &dataTemps, nil
}
