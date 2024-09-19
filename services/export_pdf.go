package services

import (
	"bytes"
	"fmt"
	"htmxll/dto"

	"github.com/jung-kurt/gofpdf"
)

func ExportPdfDaily(dailyData []dto.DataTmps, sName string, bName string) (*bytes.Buffer, error) {

	substationName := sName
	bayName := fmt.Sprintf("XXX kV %s No. XX", bName)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	imagePath := "static/css/icons/image.png"
	pdf.Image(imagePath, 10, 15, 20, 0, false, "", 0, "")
	// Set font
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(187, 10, "Daily Load Report", "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, substationName, "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, bayName, "LTRB", 0, "C", false, 0, "")
	pdf.SetY(40)

	// Header
	headers := []string{"Date", "Time", "Vab (kV)", "Vbc (kV)", "Vca (kV)", "Ia (A)", "Ib (A)", "Ic (A)", "P (MW)", "Q (MVAR)", "PF (%)"}
	for _, header := range headers {
		pdf.CellFormat(17, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 8)

	for _, data := range dailyData {
		pdf.CellFormat(17, 10, data.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, data.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseA), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseB), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseC), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.ActivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.ReactivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.PowerFactor), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}
	var buf bytes.Buffer
	err := pdf.Output(&buf)

	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func ExportPdfMonthly(day []dto.DataTmps, night []dto.DataTmps, all []dto.DataTmps, sName string, bName string) (*bytes.Buffer, error) {

	substationName := sName
	bayName := fmt.Sprintf("XXX kV %s No. XX", bName)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	imagePath := "static/css/icons/image.png"
	pdf.Image(imagePath, 10, 15, 20, 0, false, "", 0, "")
	// Set font
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(187, 10, "Daily Load Report", "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, substationName, "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, bayName, "", 0, "C", false, 0, "")
	pdf.SetY(40)

	pdf.Ln(-1)
	pdf.CellFormat(187, 10, "Day Time Peak", "LTRB", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, "08:00-15:30", "LTRB", 0, "C", false, 0, "")
	pdf.Ln(-1)

	// Header
	headers := []string{"Date", "Time", "Vab (kV)", "Vbc (kV)", "Vca (kV)", "Ia (A)", "Ib (A)", "Ic (A)", "P (MW)", "Q (MVAR)", "PF (%)"}
	for _, header := range headers {
		pdf.CellFormat(17, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 8)

	for _, data := range day {
		pdf.CellFormat(17, 10, data.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, data.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseA), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseB), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseC), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.ActivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.ReactivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.PowerFactor), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	pdf.SetFont("Arial", "B", 10)
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, "Night Time Peak", "LTRB", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, "00:00-07:30, 16:00-23:30", "LTRB", 0, "C", false, 0, "")
	pdf.Ln(-1)
	//night

	for _, header := range headers {
		pdf.CellFormat(17, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 8)

	for _, data := range night {
		pdf.CellFormat(17, 10, data.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, data.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseA), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseB), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseC), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.ActivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.ReactivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.PowerFactor), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}
	pdf.SetFont("Arial", "B", 10)
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, "Day Time Peak", "LTRB", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, "08:00-15:30", "LTRB", 0, "C", false, 0, "")
	pdf.Ln(-1)

	// Header
	for _, header := range headers {
		pdf.CellFormat(17, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 8)

	for _, data := range all {
		pdf.CellFormat(17, 10, data.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, data.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseA), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseB), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseC), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.ActivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.ReactivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.PowerFactor), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)

	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func ExportPdfYearly(peak []dto.DataTmpsYear, light []dto.DataTmpsYear, sName string, bName string) (*bytes.Buffer, error) {

	substationName := sName
	bayName := fmt.Sprintf("XXX kV %s No. XX", bName)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	imagePath := "static/css/icons/image.png"
	pdf.Image(imagePath, 10, 15, 20, 0, false, "", 0, "")
	// Set font
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(187, 10, "Daily Load Report", "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, substationName, "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, bayName, "", 0, "C", false, 0, "")
	pdf.SetY(40)

	pdf.Ln(-1)
	pdf.CellFormat(187, 10, "Peak", "", 0, "C", false, 0, "")
	pdf.Ln(-1)

	// Header
	headers := []string{"Month", "Date", "Time", "Vab (kV)", "Vbc (kV)", "Vca (kV)", "Ia (A)", "Ib (A)", "Ic (A)", "P (MW)", "Q (MVAR)", "PF (%)"}
	for _, header := range headers {
		pdf.CellFormat(16, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 8)

	for _, data := range peak {
		pdf.CellFormat(16, 10, data.Month, "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, data.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, data.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, fmt.Sprintf("%.2f", data.CurrentPhaseA), "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, fmt.Sprintf("%.2f", data.CurrentPhaseB), "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, fmt.Sprintf("%.2f", data.CurrentPhaseC), "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, fmt.Sprintf("%.2f", data.ActivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, fmt.Sprintf("%.2f", data.ReactivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, fmt.Sprintf("%.2f", data.PowerFactor), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	pdf.SetFont("Arial", "B", 10)
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, "Light Load", "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	//night

	for _, header := range headers {
		pdf.CellFormat(16, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 8)

	for _, data := range light {
		pdf.CellFormat(16, 10, data.Month, "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, data.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, data.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, fmt.Sprintf("%.2f", data.CurrentPhaseA), "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, fmt.Sprintf("%.2f", data.CurrentPhaseB), "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, fmt.Sprintf("%.2f", data.CurrentPhaseC), "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, fmt.Sprintf("%.2f", data.ActivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, fmt.Sprintf("%.2f", data.ReactivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(16, 10, fmt.Sprintf("%.2f", data.PowerFactor), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)

	if err != nil {
		return nil, err
	}
	return &buf, nil
}
