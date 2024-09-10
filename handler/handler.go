package handler

import (
	"fmt"
	"htmxll/entity"
	"htmxll/filter"
	"htmxll/services"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetOptionText(c echo.Context) error
	GetDailyReport(c echo.Context) error
}

type handler struct {
	srv services.Service
}

func NewHandler(srv services.Service) Handler {
	return handler{srv}
}

func (h handler) GetDailyReport(c echo.Context) error {
	data, err := h.srv.GetLatestData(5, filter.SortData{})
	if err != nil {
		return c.Render(200, "daily", []entity.DataTmps{})
	}
	return c.Render(200, "daily", data)
}

func (h handler) GetOptionText(c echo.Context) error {
	name := c.QueryParam("name")
	if name != "" {

		return c.String(200, fmt.Sprintf(`<span id="text-option">%s</span>`, name))
	}
	return c.String(200, ` <span id="text-option">%s</span>`)
}
