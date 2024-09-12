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
	"log"
	"os"
	"path/filepath"

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
	config := config.NewConfig()
	db := database.NewPostgresDatabase(config)
	repo := repository.NewRepository(db)
	service := services.NewService(repo)
	f := excelize.NewFile()
	defer f.Close()

	readFile := filedata.NewFileData(repo)
	go readFile.CheckNewFileRealTime()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "static")
	e.Renderer = newTemplate()

	hand := handler.NewHandler(service)

	//page := newPage()
	count := Count{Count: 0}
	e.GET("/", func(c echo.Context) error {
		//c.Render(200, "index", nil)
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

	status := false
	e.GET("/datetime-option", func(c echo.Context) error {
		status = !status
		if !status {
			return c.Render(200, "option", models.DateTimeOption{})
		}
		return c.Render(200, "option", models.NewDateTimeOption())
	})
	stationOptionStatus := false
	e.GET("/station-option", func(c echo.Context) error {
		data, err := repo.GetSubStations()
		if err != nil {
			log.Println("this error")
			return c.Render(200, "option-station", []entity.SubStation{})
		}
		stationOptionStatus = !stationOptionStatus
		if !stationOptionStatus {
			return c.Render(200, "option-station", []entity.SubStation{})
		}
		return c.Render(200, "option-station", data)
	})
	bayOptionStatus := false
	e.GET("/bay-option", func(c echo.Context) error {
		data, err := repo.GetBays()
		if err != nil {
			log.Println("this error", err)
			return c.Render(200, "option", []entity.Bay{})
		}
		bayOptionStatus = !bayOptionStatus
		if !bayOptionStatus {
			return c.Render(200, "option-bay", []entity.Bay{})
		}
		return c.Render(200, "option-bay", data)
	})

	e.GET("/text-option", hand.GetOptionText)
	e.GET("/text-station-option", hand.GetStationOptionText)
	e.GET("/text-bay-option", hand.GetBayOptionText)
	e.GET("/data", hand.GetDailyReport)
	e.Logger.Fatal(e.Start(":3000"))
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
