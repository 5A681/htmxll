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
	ttime := ""

	config := config.NewConfig()
	db := database.NewPostgresDatabase(config)
	repo := repository.NewRepository(db)
	service := services.NewService(repo)
	f := excelize.NewFile()
	defer f.Close()

	readFile := filedata.NewFileData(repo)
	go readFile.CheckNewFileRealTime()
	go readFile.InitReadFile()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "static")
	e.Renderer = newTemplate()

	excelFile := excelize.NewFile()
	hand := handler.NewHandler(service, services.NewExportExcel(excelFile), &timeSpace, &stationName, &bayName, &stationId, &bayId, &ttime)
	defaultData := models.DefaultData{
		OptionDateTime: "Optioin",
	}
	changeOptHandler := handler.NewChangeOption(&defaultData)

	//page := newPage()
	count := Count{Count: 0}
	e.GET("/", func(c echo.Context) error {
		//c.Render(200, "index", nil)
		timeSpace = ""
		stationName = ""
		bayName = ""
		stationId = 0
		bayId = 0
		return c.Render(200, "index", nil)
	})
	e.GET("/monthly", func(c echo.Context) error {
		//c.Render(200, "index", nil)
		return c.Render(200, "monthly", nil)
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

	e.Logger.Fatal(e.Start(":3000"))
}
