package handler

import (
	"fmt"
	"htmxll/filter"
	"htmxll/services"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetOptionText(c echo.Context) error
	GetDailyReport(c echo.Context) error
	GetStationOptionText(c echo.Context) error
	GetBayList(c echo.Context) error
	GetStationList(c echo.Context) error
	ExportPdf(c echo.Context) error
	ExportExcel(c echo.Context) error
}

type handler struct {
	srv   services.Service
	excel services.ExportExcel
}

func NewHandler(srv services.Service, excel services.ExportExcel) Handler {
	return handler{srv, excel}
}

var timeSpace string = "daily"
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
		timeSpace = c.QueryParam("component")
	}
	log.Println("component = ", timeSpace)

	if c.QueryParam("station") != "" {
		log.Println("station = ", c.QueryParam("station"))
		id, err := strconv.Atoi(c.QueryParam("station"))
		if err != nil {
			log.Println(err)
		} else {
			stationId = id
			s, err := h.srv.GetSubStationById(stationId)
			if err != nil {
				log.Println(err)
			} else {
				bayName = s.Name
			}

		}
	}
	if c.QueryParam("bay") != "" {
		id, err := strconv.Atoi(c.QueryParam("bay"))
		if err != nil {
			log.Println(err)
		} else {
			bayId = id
			bay, err := h.srv.GetBayById(bayId)
			if err != nil {
				log.Println(err)
			} else {
				bayName = bay.Name
			}

		}
	}

	if timeSpace != "" {
		if timeSpace == "daily" {
			data, err := h.srv.GetLatestData(bayId, filter.SortData{})
			if err != nil {
				return c.Render(200, "daily", response)
			}

			response["DailyData"] = data
			return c.Render(200, "content", response)
		} else if timeSpace == "monthly" {
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
		} else if timeSpace == "yearly" {
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
		s, err := h.srv.GetSubStationById(stationId)
		if err != nil {
			log.Println(err)
		} else {
			stationName = s.Name
		}
	}
	res, err := h.srv.GetAllBay(stationId)
	if err != nil {
		log.Println(err)
		return c.Render(200, "bay-list", nil)
	}
	data := map[string]interface{}{
		"Data": res,
		"Time": timeSpace,
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
		"Time": timeSpace,
	}
	return c.Render(200, "station-list", data)
}

func DeleteFile() {
	pdfFiles, err := filepath.Glob(filepath.Join("", "*.pdf"))
	if err != nil {
		log.Println(err)
	}

	// Iterate over the list of files and delete each one
	for _, file := range pdfFiles {
		err := os.Remove(file)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Deleted: %s\n", file)
	}
	xlsxFiles, err := filepath.Glob(filepath.Join("", "*.xlsx"))
	if err != nil {
		log.Println(err)
	}
	for _, file := range xlsxFiles {
		err := os.Remove(file)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Deleted: %s\n", file)
	}
}

func (h handler) ExportPdf(c echo.Context) error {

	if timeSpace == "daily" {
		datas, err := h.srv.GetLatestData(bayId, filter.SortData{})
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		buf, err := services.ExportPdfDaily(datas, stationName, bayName)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}

		c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=example.pdf")
		c.Response().Header().Set(echo.HeaderContentType, "application/pdf")

		// Send the PDF as a binary response
		return c.Blob(http.StatusOK, "application/pdf", buf.Bytes())
	} else if timeSpace == "monthly" {

		day, err := h.srv.GetDataLatestMonthDayTime(bayId, filter.SortData{})
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		night, err := h.srv.GetDataLatestMonthNightTime(bayId, filter.SortData{})
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		all, err := h.srv.GetDataLatestMonthAllTime(bayId, filter.SortData{})
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		buf, err := services.ExportPdfMonthly(day, night, all, stationName, bayName)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}

		c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=example.pdf")
		c.Response().Header().Set(echo.HeaderContentType, "application/pdf")

		// Send the PDF as a binary response
		return c.Blob(http.StatusOK, "application/pdf", buf.Bytes())
	}

	peak, err := h.srv.GetDataLatestYearPeakTime(bayId, 2024, filter.SortData{})
	if err != nil {
		log.Println("err:", err.Error())
		return c.String(200, ``)
	}
	light, err := h.srv.GetDataLatestYearLightTime(bayId, 2024, filter.SortData{})
	if err != nil {
		log.Println("err:", err.Error())
		return c.String(200, ``)
	}

	buf, err := services.ExportPdfYearly(peak, light, stationName, bayName)
	if err != nil {
		log.Println("err:", err.Error())
		return c.String(200, ``)
	}

	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=example.pdf")
	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")

	// Send the PDF as a binary response
	return c.Blob(http.StatusOK, "application/pdf", buf.Bytes())

}

func (h handler) ExportExcel(c echo.Context) error {
	datas, err := h.srv.GetLatestData(bayId, filter.SortData{})
	if err != nil {
		log.Println("err:", err.Error())
		return c.String(200, ``)
	}
	err = h.excel.ExportExcelDaily(datas, "test.xlsx")
	if err != nil {
		log.Println("err:", err.Error())
		return c.String(200, ``)
	}
	defer os.Remove("test.xlsx")
	return c.Attachment("test.xlsx", "example.xlsx")
}
