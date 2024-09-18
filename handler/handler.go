package handler

import (
	"fmt"
	"htmxll/filter"
	"htmxll/services"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetOptionText(c echo.Context) error
	GetDailyReport(c echo.Context) error
	GetStationOptionText(c echo.Context) error
	GetBayList(c echo.Context) error
	GetStationList(c echo.Context) error
}

type handler struct {
	srv services.Service
}

func NewHandler(srv services.Service) Handler {
	return handler{srv}
}

var time string = "daily"
var stationName string
var bayName string
var stationId int
var bayId int

func (h handler) GetDailyReport(c echo.Context) error {
	response := map[string]interface{}{
		"DailyData":   nil,
		"MonthlyData": nil,
		"YearlyData":  nil,
		"StationName": stationName,
		"BayName":     bayName,
	}
	if c.QueryParam("component") != "" {
		time = c.QueryParam("component")
	}
	log.Println("component = ", time)

	if c.QueryParam("station") != "" {
		id, err := strconv.Atoi(c.QueryParam("station"))
		if err != nil {
			log.Println(err)
		} else {
			stationId = id
		}
	}
	if c.QueryParam("bay") != "" {
		id, err := strconv.Atoi(c.QueryParam("bay"))
		if err != nil {
			log.Println(err)
		} else {
			bayId = id
		}
	}
	log.Println("bay = ", bayId)

	if time != "" {
		if time == "daily" {
			data, err := h.srv.GetLatestData(bayId, filter.SortData{})
			if err != nil {
				return c.Render(200, "daily", response)
			}

			response["DailyData"] = data
			return c.Render(200, "content", response)
		} else if time == "monthly" {
			DayData, err := h.srv.GetDataLatestMonthDayTime(bayId, filter.SortData{})
			if err != nil {
				return c.Render(200, "content", response)
			}
			NightData, err := h.srv.GetDataLatestMonthNightTime(bayId, filter.SortData{})
			if err != nil {
				return c.Render(200, "content", response)
			}
			AllData, err := h.srv.GetDataLatestMonthAllTime(bayId, filter.SortData{})
			if err != nil {
				return c.Render(200, "content", response)
			}

			response["MonthlyData"] = map[string]interface{}{"Day": DayData, "Night": NightData, "All": AllData}
			return c.Render(200, "content", response)
		} else if time == "yearly" {
			peak, err := h.srv.GetDataLatestYearPeakTime(bayId, 2024, filter.SortData{})
			if err != nil {
				return c.Render(200, "content", response)
			}
			light, err := h.srv.GetDataLatestYearLightTime(bayId, 2024, filter.SortData{})
			if err != nil {
				log.Println("error = ", err)
				return c.Render(200, "content", response)
			}
			response["YearlyData"] = map[string]interface{}{"Peak": peak, "Light": light}
			return c.Render(200, "content", response)
		}
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

func (h handler) GetBayList(c echo.Context) error {

	station, err := h.srv.GetFirstSubstation()
	if err != nil {
		log.Println(err)
		return c.Render(200, "bay-list", nil)
	}
	stationId = station.Id
	if c.QueryParam("station") != "" {
		id, err := strconv.Atoi(c.QueryParam("station"))
		if err != nil {
			log.Println(err)
			return c.Render(200, "bay-list", nil)
		}
		stationId = id
	}
	res, err := h.srv.GetAllBay(stationId)
	if err != nil {
		log.Println(err)
		return c.Render(200, "bay-list", nil)
	}
	data := map[string]interface{}{
		"Data": res,
		"Time": time,
	}
	return c.Render(200, "bay-list", data)
}
func (h handler) GetStationList(c echo.Context) error {
	res, err := h.srv.GetAllSubStation()
	if err != nil {
		return c.Render(200, "station-list", nil)
	}
	data := map[string]interface{}{
		"Data": res,
		"Time": time,
	}
	return c.Render(200, "station-list", data)
}
