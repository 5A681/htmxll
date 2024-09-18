package handler

import (
	"htmxll/models"

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
