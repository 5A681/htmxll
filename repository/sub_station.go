package repository

import (
	"htmxll/entity"
	"time"
)

func (s repository) GetSubStationById(id int) (*entity.SubStation, error) {
	var subStation entity.SubStation
	err := s.db.Get(&subStation, `select * from sub_stations where id = $1 order by id asc`, id)
	if err != nil {
		return nil, err
	}
	return &subStation, nil
}
func (s repository) GetFirstSubstation() (*entity.SubStation, error) {
	var subStation entity.SubStation
	err := s.db.Get(&subStation, `select * from sub_stations  order by id asc`)
	if err != nil {
		return nil, err
	}
	return &subStation, nil
}

func (s repository) GetSubStations() ([]entity.SubStation, error) {
	var subStations []entity.SubStation
	err := s.db.Select(&subStations, "SELECT * FROM sub_stations ORDER BY created_at ASC")
	if err != nil {
		return nil, err
	}
	return subStations, nil
}

func (s repository) GetSubStationByName(name string) (*entity.SubStation, error) {
	var subStation entity.SubStation
	err := s.db.Get(&subStation, `select * from sub_stations where name = $1 order by name asc`, name)
	if err != nil {
		return nil, err
	}
	return &subStation, nil
}

func (s repository) CreateSubStation(sub *entity.SubStation) error {
	sqlCreate := `insert into sub_stations(name,created_at) values($1,$2)`
	_, err := s.db.Exec(sqlCreate, sub.Name, time.Now())
	if err != nil {
		return err
	}
	return nil
}
