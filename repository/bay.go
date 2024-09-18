package repository

import (
	"htmxll/entity"
	"log"
	"time"
)

func (s repository) GetBayById(id int) (*entity.Bay, error) {
	var bay entity.Bay
	err := s.db.Get(&bay, `select * from bays where id = $1 order by id asc`, id)
	if err != nil {
		return nil, err
	}
	return &bay, nil
}
func (s repository) GetBays(stationId int) ([]entity.Bay, error) {

	var bays []entity.Bay
	err := s.db.Select(&bays, `select * from bays where sub_station_id = $1 order by CAST(SUBSTRING(name FROM '[0-9]+') AS INTEGER) ASC;`, stationId)
	if err != nil {
		return nil, err
	}
	return bays, nil
}

func (s repository) GetBayByNameAndSubStationId(id int, name string) (*entity.Bay, error) {
	log.Println(id, name)
	var bay entity.Bay
	err := s.db.Get(&bay, `select * from bays where sub_station_id = $1 and name = $2 order by id asc limit 1`, id, name)
	if err != nil {
		return nil, err
	}
	return &bay, nil
}

func (s repository) CreateBay(bay *entity.Bay) error {
	sqlCreate := `insert into bays(name,created_at,sub_station_id) values($1,$2,$3)`
	_, err := s.db.Exec(sqlCreate, bay.Name, time.Now(), bay.SubStationId)
	if err != nil {
		return err
	}
	return nil
}
