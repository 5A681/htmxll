package models

type DateTime struct {
	Id     int
	Option string
	Page   string
}

type DateTimeOption struct {
	DateTimes []DateTime
}

func NewDateTimeOption() DateTimeOption {
	return DateTimeOption{
		[]DateTime{
			{1, "Daily", "daily"},
			{2, "Monthly", "monthly"},
			{3, "Yearly", "yearly"},
		},
	}
}
