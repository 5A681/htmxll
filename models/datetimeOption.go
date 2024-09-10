package models

type DateTime struct {
	Id     int
	Option string
}

type DateTimeOption struct {
	DateTimes []DateTime
}

func NewDateTimeOption() DateTimeOption {
	return DateTimeOption{
		[]DateTime{
			{1, "Daily"},
			{2, "Month"},
			{3, "Year"},
		},
	}
}
