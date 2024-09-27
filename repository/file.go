package repository

import (
	"htmxll/entity"
	"time"
)

func (r repository) GetFileName(name string) (*entity.FileTemps, error) {
	var subStation entity.FileTemps
	err := r.db.Get(&subStation, `select * from file_temps where dir_name = $1 `, name)
	if err != nil {
		return nil, err
	}
	return &subStation, nil
}

func (r repository) CreateFileTemps(file *entity.FileTemps) error {
	sqlCreate := `insert into file_temps(dir_name,created_at) values($1,$2)`
	_, err := r.db.Exec(sqlCreate, file.DirName, time.Now())
	if err != nil {
		return err
	}
	return nil
}
