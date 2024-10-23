package dto

type MonthlyRowData struct {
	Bay       string
	PeakDay   MonthlyData
	PeakNight MonthlyData
	AllLow    MonthlyData
}

type MonthlyData struct {
	Date string
	Time string
	Kv   float32
	Ia   float32
	Ib   float32
	Ic   float32
	Mw   float32
	Mvar float32
}
