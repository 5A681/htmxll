package handler

import (
	"fmt"
	"htmxll/filter"
	"htmxll/services"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetOptionText(c echo.Context) error
	GetDailyReport(c echo.Context) error
	GetStationOptionText(c echo.Context) error
	GetBayOptionText(c echo.Context) error
}

type handler struct {
	srv services.Service
}

func NewHandler(srv services.Service) Handler {
	return handler{srv}
}

func (h handler) GetDailyReport(c echo.Context) error {
	response := map[string]interface{}{
		"DailyData":   nil,
		"MonthlyData": nil,
		"YearlyData":  nil,
	}
	if c.QueryParam("component") == "daily" {

		data, err := h.srv.GetLatestData(5, filter.SortData{})
		if err != nil {
			return c.Render(200, "daily", response)
		}
		response["DailyData"] = data
		return c.Render(200, "content", response)
	} else if c.QueryParam("component") == "monthly" {
		DayData, err := h.srv.GetDataLatestMonthDayTime(5, filter.SortData{})
		if err != nil {
			return c.Render(200, "content", response)
		}
		NightData, err := h.srv.GetDataLatestMonthNightTime(5, filter.SortData{})
		if err != nil {
			return c.Render(200, "content", response)
		}
		AllData, err := h.srv.GetDataLatestMonthAllTime(5, filter.SortData{})
		if err != nil {
			return c.Render(200, "content", response)
		}

		response["MonthlyData"] = map[string]interface{}{"Day": DayData, "Night": NightData, "All": AllData}
		return c.Render(200, "content", response)
	}
	return c.Render(200, "content", response)

}

func (h handler) GetOptionText(c echo.Context) error {
	name := c.QueryParam("name")
	if name != "" {

		return c.String(200, fmt.Sprintf(`<span id="text-option">%s</span>`, name))
	}
	return c.String(200, ` <span id="text-option">Options</span>`)
}

func (h handler) GetStationOptionText(c echo.Context) error {
	name := c.QueryParam("name")
	if name != "" {

		return c.String(200, fmt.Sprintf(`<span id="text-station-option">%s</span>`, name))
	}
	return c.String(200, ` <span id="text-station-option">Stations</span>`)
}

func (h handler) GetBayOptionText(c echo.Context) error {
	name := c.QueryParam("name")
	if name != "" {

		return c.String(200, fmt.Sprintf(`<span id="text-bay-option">%s</span>`, name))
	}
	return c.String(200, ` <span id="text-bay-option">Bays</span>`)
}
