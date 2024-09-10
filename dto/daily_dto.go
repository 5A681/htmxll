package dto

import "time"

type DataTmps struct {
	Id            int
	CurrentPhaseA float32
	CurrentPhaseB float32
	CurrentPhaseC float32
	ActivePower   float32
	ReactivePower float32
	PowerFactor   float32
	Date          string
	Time          string
	CreatedAt     time.Time
	BayId         int
}
