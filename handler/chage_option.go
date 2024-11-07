package handler

import (
	"fmt"
	"htmxll/models"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

type ChangeOption interface {
	GetOptionDateTimmeText(c echo.Context) error
}

type changeOption struct {
	defaultData *models.DefaultData
}

func NewChangeOption(defaultData *models.DefaultData) ChangeOption {
	return &changeOption{defaultData}
}

func (h *changeOption) GetOptionDateTimmeText(c echo.Context) error {
	type Data struct {
		Name string
	}
	d := Data{Name: "Daily"}
	return c.Render(200, "dropdown-datetime", d)
}

func (h handler) GetDateTimePickerFormat(c echo.Context) error {
	defaultTime := ""
	defaultMonth := ""

	t, err := time.Parse("02/01/2006", *h.time)
	if err != nil {
		log.Println("error defaultTime", defaultTime)
	} else {
		defaultTime = t.Format("2006-01-02")
		defaultMonth = t.Format("2006-01")
	}
	log.Println("defualt time", defaultTime, "oldtime", *h.time)
	if *h.timeSpace == "daily" || *h.timeSpace == "" {
		return c.String(200, fmt.Sprintf(` <div id="date-picker-input">
                    <input id="datepicker" name="time" type="date" hx-get="/data" hx-trigger="change" hx-target="#content" hx-swap="innerHTML" 
                    hx-include="#datepicker" value="%s"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full ps-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    placeholder="Select date">
                </div>`, defaultTime))
	} else if *h.timeSpace == "monthly" {
		return c.String(200, fmt.Sprintf(` <div id="date-picker-input">
                    <input id="datepicker" name="time" type="month" hx-get="/data" hx-trigger="change" hx-target="#content" hx-swap="innerHTML" 
                    hx-include="#datepicker" value="%s"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full ps-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    placeholder="Select date">
                </div>`, defaultMonth))
	}
	return c.String(200, `<div id="date-picker-input">
                    <input id="datepicker" name="time" type="date" hx-get="/data" hx-trigger="change" hx-target="#content" hx-swap="innerHTML" 
                    hx-include="#datepicker"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full ps-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    placeholder="Select date">
                </div>`)
}
