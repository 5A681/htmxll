package services

import (
	"bytes"
	"fmt"
	"htmxll/dto"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func ExportPdfDaily(dailyData []dto.DataTmps, sName string, bName string, exportHeader string) (*bytes.Buffer, error) {
	bayName := bName
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	imagePath := "static/css/icons/image.png"
	pdf.Image(imagePath, 10, 15, 20, 0, false, "", 0, "")
	// Set font
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(189, 10, "Daily Load Report", "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(189, 10, exportHeader, "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(189, 10, bayName, "LTRB", 0, "C", false, 0, "")
	pdf.SetY(40)

	// Header
	headers := []string{"Date", "Time", "kV", "Ia", "Ib", "Ic", "MW", "MVAR", "PF"}
	for _, header := range headers {
		pdf.CellFormat(21, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 8)

	for i, data := range dailyData {
		if i == 22 {
			pdf.SetFont("Arial", "B", 10)
			headers := []string{"Date", "Time", "kV", "Ia", "Ib", "Ic", "MW", "MVAR", "PF"}
			for _, header := range headers {
				pdf.CellFormat(21, 10, header, "1", 0, "C", false, 0, "")
			}
			pdf.SetFont("Arial", "", 8)
			pdf.Ln(-1)
		}
		pdf.CellFormat(21, 10, data.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, data.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.Kv), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.CurrentPhaseA), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.CurrentPhaseB), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.CurrentPhaseC), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.ActivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.ReactivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.PowerFactor), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}
	var buf bytes.Buffer
	err := pdf.Output(&buf)

	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func ExportPdfMonthly(items []dto.MonthlyRowData, sName string, bName string, exportHeader string) (*bytes.Buffer, error) {
	log.Println("Substation name = ", exportHeader)
	substationName := exportHeader
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	imagePath := "static/css/icons/image.png"
	pdf.Image(imagePath, 10, 5, 20, 0, false, "", 0, "")
	// Set font
	pdf.SetFont("Arial", "B", 6)
	pdf.CellFormat(270, 5, "Monthly Load Report", "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(270, 5, substationName, "", 0, "C", false, 0, "")
	// pdf.SetY(35)

	pdf.Ln(-1)
	pdf.CellFormat(100, 10, "ALL DAY PEAK LOAD", "LTRB", 0, "C", false, 0, "")
	pdf.CellFormat(85, 10, "ALL NIGHT LOW LOAD", "LTRB", 0, "C", false, 0, "")
	pdf.CellFormat(85, 10, "FULL DAY LOW LOAD", "LTRB", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(100, 10, "08:00-15:30", "LTRB", 0, "C", false, 0, "")
	pdf.CellFormat(85, 10, "16:00-23:30", "LTRB", 0, "C", false, 0, "")
	pdf.CellFormat(85, 10, "00:00-23:30", "LTRB", 0, "C", false, 0, "")
	pdf.Ln(-1)

	// Header
	pdf.CellFormat(15, 8, "TP1/25MVA", "1", 0, "C", false, 0, "")
	headers := []string{"Date", "Time", "kV", "Ia", "Ib", "Ic", "MW", "MVAR"}
	for i, header := range headers {
		if i == 0 {
			pdf.CellFormat(15, 8, header, "1", 0, "C", false, 0, "")
		} else {
			pdf.CellFormat(10, 8, header, "1", 0, "C", false, 0, "")
		}
	}
	for i, header := range headers {
		if i == 0 {
			pdf.CellFormat(15, 8, header, "1", 0, "C", false, 0, "")
		} else {
			pdf.CellFormat(10, 8, header, "1", 0, "C", false, 0, "")
		}
	}
	for i, header := range headers {
		if i == 0 {
			pdf.CellFormat(15, 8, header, "1", 0, "C", false, 0, "")
		} else {
			pdf.CellFormat(10, 8, header, "1", 0, "C", false, 0, "")
		}
	}

	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 5)

	for _, data := range items {
		pdf.CellFormat(15, 8, data.Bay, "1", 0, "C", false, 0, "")
		pdf.CellFormat(15, 8, data.PeakDay.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, data.PeakDay.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.PeakDay.Kv), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.PeakDay.Ia), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.PeakDay.Ib), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.PeakDay.Ic), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.PeakDay.Mw), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.PeakDay.Mvar), "1", 0, "C", false, 0, "")

		pdf.CellFormat(15, 8, data.PeakNight.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, data.PeakNight.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.PeakNight.Kv), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.PeakNight.Ia), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.PeakNight.Ib), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.PeakNight.Ic), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.PeakNight.Mw), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.PeakNight.Mvar), "1", 0, "C", false, 0, "")

		pdf.CellFormat(15, 8, data.AllLow.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, data.AllLow.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.AllLow.Kv), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.AllLow.Ia), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.AllLow.Ib), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.AllLow.Ic), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.AllLow.Mw), "1", 0, "C", false, 0, "")
		pdf.CellFormat(10, 8, fmt.Sprintf("%.2f", data.AllLow.Mvar), "1", 0, "C", false, 0, "")

		pdf.Ln(-1)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)

	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func ExportPdfYearly(peak []dto.DataTmpsYear, light []dto.DataTmpsYear, sName string, bName string, exportHeader string) (*bytes.Buffer, error) {

	substationName := exportHeader
	bayName := bName
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	imagePath := "static/css/icons/image.png"
	pdf.Image(imagePath, 10, 15, 20, 0, false, "", 0, "")
	// Set font
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(189, 10, "Yearly Load Report", "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(189, 10, substationName, "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(189, 10, bayName, "", 0, "C", false, 0, "")
	pdf.SetY(40)

	pdf.Ln(-1)
	pdf.CellFormat(189, 10, "Peak Load", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	// Header
	headers := []string{"Month", "Date", "Time", "Vbc (kV)", "Ia (A)", "Ib (A)", "Ic (A)", "P (MW)", "Q (MVAR)"}
	for _, header := range headers {
		pdf.CellFormat(21, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 8)

	for _, data := range peak {
		pdf.CellFormat(21, 10, data.Month, "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, data.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, data.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.Kv), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.CurrentPhaseA), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.CurrentPhaseB), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.CurrentPhaseC), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.ActivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.ReactivePower), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	pdf.SetFont("Arial", "B", 10)
	pdf.Ln(-1)
	pdf.CellFormat(189, 10, "Light Load", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)
	//night

	for _, header := range headers {
		pdf.CellFormat(21, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 8)

	for _, data := range light {
		pdf.CellFormat(21, 10, data.Month, "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, data.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, data.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.Kv), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.CurrentPhaseA), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.CurrentPhaseB), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.CurrentPhaseC), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.ActivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(21, 10, fmt.Sprintf("%.2f", data.ReactivePower), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)

	if err != nil {
		return nil, err
	}
	return &buf, nil
}
