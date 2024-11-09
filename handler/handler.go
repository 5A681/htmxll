package handler

import (
	"fmt"
	"htmxll/config"
	"htmxll/filter"
	"htmxll/services"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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
	SelectDate(c echo.Context) error
	GetMonthBayList(c echo.Context) error
	GetMonthlyDay(c echo.Context) error
	GetMonthlyNight(c echo.Context) error
	GetMonthlyAll(c echo.Context) error
	GetRowsMonthlyData(c echo.Context) error
	GetDateTimePickerFormat(c echo.Context) error
}

type handler struct {
	srv         services.Service
	excel       services.ExportExcel
	timeSpace   *string
	stationName *string
	bayName     *string
	stationId   *int
	bayId       *int
	time        *time.Time
	month       *int
	year        *int
	day         *int
	config      config.Config
}

func NewHandler(srv services.Service, excel services.ExportExcel, timeSpace *string,
	stationName *string,
	bayName *string,
	stationId *int,
	bayId *int,
	time *time.Time, month *int, year *int, day *int, config config.Config) Handler {

	return handler{srv, excel, timeSpace, stationName, bayName, stationId, bayId, time, month, year, day, config}
}

func (h handler) GetDailyReport(c echo.Context) error {
	response := map[string]interface{}{
		"DailyData":   nil,
		"MonthlyData": nil,
		"YearlyData":  nil,
		"StationName": h.stationName,
		"BayName":     h.bayName,
	}
	if c.QueryParam("component") != "" {
		*h.timeSpace = c.QueryParam("component")
	}
	log.Println("component = ", *h.timeSpace)

	if c.QueryParam("station") != "" {
		log.Println("station = ", c.QueryParam("station"))
		id, err := strconv.Atoi(c.QueryParam("station"))
		if err != nil {
			log.Println(err)
		} else {
			*h.stationId = id
			s, err := h.srv.GetSubStationById(*h.stationId)
			if err != nil {
				log.Println(err)
			} else {
				*h.bayName = s.Name
			}

		}
	}
	if c.QueryParam("bay") != "" {
		id, err := strconv.Atoi(c.QueryParam("bay"))
		if err != nil {
			log.Println(err)
		} else {
			*h.bayId = id
			bay, err := h.srv.GetBayById(*h.bayId)
			if err != nil {
				log.Println(err)
			} else {
				*h.bayName = bay.Name
			}

		}
	}
	if c.QueryParam("time") != "" {
		var err error
		t := c.QueryParam("time")
		if len(strings.Split(t, "-")) == 3 {
			*h.time, err = time.Parse("2006-01-02", t)
			if err != nil {
				log.Println("error time day", err, *h.time)
			}
		} else if len(strings.Split(t, "-")) == 2 {
			t += "-01"
			log.Println("Phongphat month", t)
			*h.time, err = time.Parse("2006-01-02", t)
			if err != nil {
				log.Println("error time month", err, *h.time)
			}
		} else {
			t += "-01-01"
			*h.time, err = time.Parse("2006-01-02", t)
			if err != nil {
				log.Println("error time year", err, *h.time)
			}
		}
	}

	if *h.timeSpace != "" {
		if *h.timeSpace == "daily" {
			data, err := h.srv.GetLatestData(*h.bayId, *h.time)
			if err != nil {
				log.Println("This Error", err)
				return c.Render(200, "daily", response)
			}
			if len(data) > 0 {
				*h.time, err = time.Parse("02/01/2006", data[0].Date)
				if err != nil {
					log.Println("error time", err, *h.time)
				}
			}

			response["DailyData"] = data
			return c.Render(200, "daily", response)
		} else if *h.timeSpace == "monthly" {
			return h.GetRowsMonthlyData(c)
		} else if *h.timeSpace == "yearly" {
			peak, err := h.srv.GetDataLatestYearPeakTime(*h.time, *h.bayId, 2024, filter.SortData{})
			if err != nil {
				return c.Render(200, "yearly", response)
			}
			light, err := h.srv.GetDataLatestYearLightTime(*h.time, *h.bayId, 2024, filter.SortData{})
			if err != nil {
				log.Println("error = ", err)
				return c.Render(200, "yearly", response)
			}
			response["YearlyData"] = map[string]interface{}{"Peak": peak, "Light": light}
			return c.Render(200, "yearly", response)
		}
	}
	return c.Render(200, "daily", response)

}
func (h handler) GetMonthBayList(c echo.Context) error {
	bays, err := h.srv.GetAllBay()
	if err != nil {
		log.Println(err)
		return c.Render(200, "first-column", bays)
	}
	return c.Render(200, "first-column", bays)
}

func (h handler) GetMonthlyDay(c echo.Context) error {
	Data, err := h.srv.GetDataLatestMonthDayTime(*h.time, *h.bayId, filter.SortData{})
	if err != nil {
		return c.Render(200, "monthly-table-day", Data)
	}
	return c.Render(200, "monthly-table-day", Data)

}
func (h handler) GetMonthlyNight(c echo.Context) error {
	Data, err := h.srv.GetDataLatestMonthNightTime(*h.time, *h.bayId, filter.SortData{})
	if err != nil {
		return c.Render(200, "monthly-table-night", Data)
	}
	return c.Render(200, "monthly-table-night", Data)
}
func (h handler) GetMonthlyAll(c echo.Context) error {
	Data, err := h.srv.GetDataLatestMonthAllTime(*h.time, *h.bayId, filter.SortData{})
	if err != nil {
		return c.Render(200, "monthly-table-all", Data)
	}
	return c.Render(200, "monthly-table-all", Data)
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

	if c.QueryParam("station") != "" {
		log.Println("hello station")
		id, err := strconv.Atoi(c.QueryParam("station"))
		if err != nil {
			data := map[string]interface{}{
				"Data": nil,
				"Name": *h.bayName,
				"Time": *h.timeSpace,
			}
			return c.Render(200, "bay-list", data)
		}
		*h.stationId = id
		s, err := h.srv.GetSubStationById(*h.stationId)
		if err != nil {
			log.Println(err)
		} else {
			*h.stationName = s.Name
		}
		log.Println("station data = ", s.Name)
	}
	if *h.timeSpace == "monthly" {
		return c.String(200, `<select 
    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-[150px] p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 " disabled
    hx-get="/data" hx-target="#content" name="bay" hx-trigger="change" hx-swap="innerHTML">
<option selected >Choose Bay</option>
   
</select>`)
	}

	res, err := h.srv.GetAllBayByStationId(*h.stationId)
	if err != nil {
		log.Println(err)
		data := map[string]interface{}{
			"Data": res,
			"Name": *h.bayName,
			"Time": *h.timeSpace,
		}
		return c.Render(200, "bay-list", data)
	}
	if *h.bayName == "" {
		*h.bayName = "Choose Bays"
	}
	data := map[string]interface{}{
		"Data": res,
		"Name": *h.bayName,
		"Time": *h.timeSpace,
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
		"Time": *h.timeSpace,
	}
	return c.Render(200, "station-list", data)
}

func (h handler) GetRowsMonthlyData(c echo.Context) error {
	res, err := h.srv.GetRowsMonthlyData(*h.time)
	if err != nil {

		return c.Render(200, "monthly-rows", res)
	}
	if len(res) > 0 {
		dataTime, err := time.Parse("02-01-2006", res[0].PeakDay.Date)
		if err != nil {
			log.Println("error time this", err, *h.time)
			dataTime, err = time.Parse("02/01/2006", res[0].PeakDay.Date)
			log.Println("error time this", err, *h.time)
		}
		if !dataTime.IsZero() {
			*h.time = dataTime
		}
	}
	return c.Render(200, "monthly-rows", res)
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

	if *h.timeSpace == "daily" {
		datas, err := h.srv.GetLatestData(*h.bayId, *h.time)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		buf, err := services.ExportPdfDaily(datas, *h.stationName, *h.bayName, h.config.EXPORT_HEADER)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}

		fileName := fmt.Sprintf("Daily-%s.pdf", time.Now().Format("2006-01-02"))

		c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", fileName))
		c.Response().Header().Set(echo.HeaderContentType, "application/pdf")

		// Send the PDF as a binary response
		return c.Blob(http.StatusOK, "application/pdf", buf.Bytes())
	} else if *h.timeSpace == "monthly" {

		data, err := h.srv.GetRowsMonthlyData(*h.time)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		buf, err := services.ExportPdfMonthly(data, *h.stationName, *h.bayName, h.config.EXPORT_HEADER)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}

		fileName := fmt.Sprintf("Monthly-%s.pdf", time.Now().Format("2006-01-02"))

		c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", fileName))
		c.Response().Header().Set(echo.HeaderContentType, "application/pdf")

		// Send the PDF as a binary response
		return c.Blob(http.StatusOK, "application/pdf", buf.Bytes())
	}

	peak, err := h.srv.GetDataLatestYearPeakTime(*h.time, *h.bayId, 2024, filter.SortData{})
	if err != nil {
		log.Println("err:", err.Error())
		return c.String(200, ``)
	}
	light, err := h.srv.GetDataLatestYearLightTime(*h.time, *h.bayId, 2024, filter.SortData{})
	if err != nil {
		log.Println("err:", err.Error())
		return c.String(200, ``)
	}

	buf, err := services.ExportPdfYearly(peak, light, *h.stationName, *h.bayName, h.config.EXPORT_HEADER)
	if err != nil {
		log.Println("err:", err.Error())
		return c.String(200, ``)
	}
	fileName := fmt.Sprintf("Yearly-%s.pdf", time.Now().Format("2006-01-02"))

	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", fileName))
	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")

	// Send the PDF as a binary response
	return c.Blob(http.StatusOK, "application/pdf", buf.Bytes())

}

func (h handler) ExportExcel(c echo.Context) error {

	if *h.timeSpace == "daily" {
		datas, err := h.srv.GetLatestData(*h.bayId, *h.time)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		err = h.excel.ExportExcelDaily(datas, "test.xlsx", *h.bayName, h.config.EXPORT_HEADER)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		fileName := fmt.Sprintf("Daily-%s.xlsx", time.Now().Format("2006-01-02"))
		defer os.Remove("test.xlsx")
		return c.Attachment("test.xlsx", fileName)
	} else if *h.timeSpace == "monthly" {
		data, err := h.srv.GetRowsMonthlyData(*h.time)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}

		err = h.excel.ExportExcelMonthly(data, "test.xlsx", *h.stationName, *h.bayName, h.config.EXPORT_HEADER)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		fileName := fmt.Sprintf("Monthly-%s.xlsx", time.Now().Format("2006-01-02"))
		defer os.Remove("test.xlsx")
		return c.Attachment("test.xlsx", fileName)
	}
	peak, err := h.srv.GetDataLatestYearPeakTime(*h.time, *h.bayId, 2024, filter.SortData{})
	if err != nil {
		log.Println("err:", err.Error())
		return c.String(200, ``)
	}
	light, err := h.srv.GetDataLatestYearLightTime(*h.time, *h.bayId, 2024, filter.SortData{})
	if err != nil {
		log.Println("err:", err.Error())
		return c.String(200, ``)
	}

	err = h.excel.ExportExcelYearly(peak, light, "test.xlsx", *h.stationName, *h.bayName)
	if err != nil {
		log.Println("err:", err.Error())
		return c.String(200, ``)
	}
	fileName := fmt.Sprintf("Yearly-%s.xlsx", time.Now().Format("2006-01-02"))
	defer os.Remove("test.xlsx")
	return c.Attachment("test.xlsx", fileName)

}

func (h handler) SelectDate(c echo.Context) error {

	if *h.timeSpace == "daily" {
		return c.String(200, `<div class="relative max-w-sm" id="select-date">
                <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                    <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" aria-hidden="true"
                        xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
                        <path
                            d="M20 4a2 2 0 0 0-2-2h-2V1a1 1 0 0 0-2 0v1h-3V1a1 1 0 0 0-2 0v1H6V1a1 1 0 0 0-2 0v1H2a2 2 0 0 0-2 2v2h20V4ZM0 18a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8H0v10Zm5-8h10a1 1 0 0 1 0 2H5a1 1 0 0 1 0-2Z" />
                    </svg>
                </div>
                <input id="datepicker" name="time" type="date" hx-get="/data" hx-trigger="change" hx-target="#content"
                    hx-include="#datepicker"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full ps-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    placeholder="Select date">
            </div>`)
	} else if *h.timeSpace == "monthly" {
		return c.String(200, `<div class="relative max-w-sm" id="select-date">
                <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                    <svg class="w-4 h-4 text-gray-200 dark:text-gray-400" aria-hidden="true"
                        xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
                        <path
                            d="M20 4a2 2 0 0 0-2-2h-2V1a1 1 0 0 0-2 0v1h-3V1a1 1 0 0 0-2 0v1H6V1a1 1 0 0 0-2 0v1H2a2 2 0 0 0-2 2v2h20V4ZM0 18a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8H0v10Zm5-8h10a1 1 0 0 1 0 2H5a1 1 0 0 1 0-2Z" />
                    </svg>
                </div>
                <input id="datepicker" name="time" type="date" hx-get="/data" hx-trigger="change" hx-target="#content" disabled
                    hx-include="#datepicker"
                    class="bg-gray-50 border border-gray-300 text-gray-200 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full ps-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    placeholder="Select date">
            </div>`)
	}
	return c.String(200, `<div class="relative">
  <button class="w-full py-2 px-4 bg-gray-300 text-gray-700 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
    Select Year
  </button>
  <div class="absolute mt-2 w-full bg-white border border-gray-300 rounded-md shadow-lg">
    <ul class="list-none p-2 space-y-1">
      <li><button class="w-full py-2 px-4 text-left text-gray-700 hover:bg-gray-200">2020</button></li>
      <li><button class="w-full py-2 px-4 text-left text-gray-700 hover:bg-gray-200">2021</button></li>
      <li><button class="w-full py-2 px-4 text-left text-gray-700 hover:bg-gray-200">2022</button></li>
      <li><button class="w-full py-2 px-4 text-left text-gray-700 hover:bg-gray-200">2023</button></li>
      <li><button class="w-full py-2 px-4 text-left text-gray-700 hover:bg-gray-200">2024</button></li>
    </ul>
  </div>
</div>`)
}
