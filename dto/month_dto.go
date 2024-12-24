package dto

type MonthlyRowData struct {
	Bay       string
	PeakDay   MonthlyData
	PeakNight MonthlyData
	All       MonthlyData
}

type MonthlyData struct {
	Date string
	Time string
	Vab  string
	Vbc  string
	Vca  string
	Ia   string
	Ib   string
	Ic   string
	Mw   string
	Mvar string
	P    string
}
