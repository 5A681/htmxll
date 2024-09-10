package repository

import (
	"htmxll/entity"
	"log"
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

func (s repository) GetSubStationByName(name string) (*entity.SubStation, error) {
	var subStation entity.SubStation
	err := s.db.Get(&subStation, `select * from sub_stations where name = $1 order by name asc`, name)
	if err != nil {
		return nil, err
	}
	return &subStation, nil
}

func (s repository) CreateSubStation(sub *entity.SubStation) error {
	log.Println("sub = ", *sub)
	sqlCreate := `insert into sub_stations(name,created_at) values($1,$2)`
	_, err := s.db.Exec(sqlCreate, sub.Name, time.Now())
	if err != nil {
		return err
	}
	return nil
}
