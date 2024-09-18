package entity

import "time"

type FileTemps struct {
	Id        int       `db:"id"`
	DirName   string    `db:"dir_name"`
	CreatedAt time.Time `db:"created_at"`
}
