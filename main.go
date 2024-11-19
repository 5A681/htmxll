package main

import (
	"html/template"
	"htmxll/config"
	"htmxll/entity"
	filedata "htmxll/file_data"
	"htmxll/handler"
	"htmxll/models"
	"htmxll/pkg/database"
	"htmxll/repository"
	"htmxll/services"
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xuri/excelize/v2"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Count struct {
	Count float32
}

type Contact struct {
	Email string
	Name  string
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func main() {

	timeSpace := ""
	stationName := ""
	bayName := ""
	stationId := 0
	bayId := 0
	ttime := time.Time{}
	year := 0
	month := 0
	day := 0

	config := config.NewConfig()
	db := database.NewPostgresDatabase(config)
	repo := repository.NewRepository(db)
	service := services.NewService(repo, config)

	readFile := filedata.NewFileData(repo)
	go readFile.CheckNewFileRealTime()
	go readFile.InitReadFile()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "static")
	e.Renderer = newTemplate()

	excelFile := excelize.NewFile()
	hand := handler.NewHandler(service, services.NewExportExcel(excelFile), &timeSpace, &stationName, &bayName, &stationId, &bayId, &ttime, &month, &year, &day, config)
	defaultData := models.DefaultData{
		OptionDateTime: "Optioin",
	}
	changeOptHandler := handler.NewChangeOption(&defaultData)
	//hello
	//page := newPage()
	count := Count{Count: 0}
	e.GET("/", func(c echo.Context) error {
		//c.Render(200, "index", nil)
		timeSpace = ""
		stationName = ""
		bayName = ""
		stationId = 0
		bayId = 0
		ttime = time.Time{}
		return c.Render(200, "index", nil)
	})
	e.GET("/monthly", func(c echo.Context) error {
		//c.Render(200, "index", nil)

		return c.Render(200, "new-monthly", nil)
	})
	e.POST("/count", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "count", count)
	})

	e.GET("/option-title", changeOptHandler.GetOptionDateTimmeText)

	// status := false
	// e.GET("/datetime-option", func(c echo.Context) error {
	// 	status = !status
	// 	if !status {
	// 		return c.Render(200, "option", models.DateTimeOption{})
	// 	}

	// 	return c.Render(200, "option", models.NewDateTimeOption())
	// })
	stationOptionStatus := false
	e.GET("/station-option", func(c echo.Context) error {
		data, err := repo.GetSubStations()
		if err != nil {
			return c.Render(200, "option-station", []entity.SubStation{})
		}
		stationOptionStatus = !stationOptionStatus
		if !stationOptionStatus {
			return c.Render(200, "option-station", []entity.SubStation{})
		}
		return c.Render(200, "option-station", data)
	})

	e.GET("/text-option", hand.GetOptionText)
	e.GET("/text-station-option", hand.GetStationOptionText)
	e.GET("/data", hand.GetDailyReport)
	e.GET("/bay-list", hand.GetBayList)
	e.GET("/station-list", hand.GetStationList)
	e.GET("/export-pdf", hand.ExportPdf)
	e.GET("/export-excel", hand.ExportExcel)
	e.GET("/select-date", hand.SelectDate)
	e.GET("/month-bay-list", hand.GetMonthBayList)
	e.GET("/month-table-day", hand.GetMonthlyDay)
	e.GET("/month-table-night", hand.GetMonthlyNight)
	e.GET("/month-table-all", hand.GetMonthlyAll)
	e.GET("/monthly-rows", hand.GetRowsMonthlyData)
	e.GET("/time-picker", hand.GetDateTimePickerFormat)
	e.GET("/test-time", func(c echo.Context) error {

		return c.String(200, `<div class="relative max-w-sm">
  <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
     <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
        <path d="M20 4a2 2 0 0 0-2-2h-2V1a1 1 0 0 0-2 0v1h-3V1a1 1 0 0 0-2 0v1H6V1a1 1 0 0 0-2 0v1H2a2 2 0 0 0-2 2v2h20V4ZM0 18a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8H0v10Zm5-8h10a1 1 0 0 1 0 2H5a1 1 0 0 1 0-2Z"/>
      </svg>
  </div>
  <input id="datepicker-format" datepicker datepicker-format="mm-dd-yyyy" type="text" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full ps-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Select date">
</div>`)

	})

	e.Logger.Fatal(e.Start(":3000"))
}
