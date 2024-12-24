package dto

import "time"

type DataTmps struct {
	Id            int
	CurrentPhaseA string
	CurrentPhaseB string
	CurrentPhaseC string
	ActivePower   string
	Vab           string
	Vbc           string
	Vca           string
	ReactivePower string
	PowerFactor   string
	Date          string
	Time          string
	CreatedAt     time.Time
	BayId         int
}

type DataTmpsYear struct {
	Id            int
	CurrentPhaseA string
	CurrentPhaseB string
	CurrentPhaseC string
	Vab           string
	Vbc           string
	Vca           string
	ActivePower   string
	ReactivePower string
	PowerFactor   string
	Month         string
	Date          string
	Time          string
	CreatedAt     time.Time
	BayId         int
}
