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
	pdf.CellFormat(198, 10, "Daily Load Report", "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(198, 10, exportHeader, "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(198, 10, bayName, "LTRB", 0, "C", false, 0, "")
	pdf.SetY(40)

	// Header
	headers := []string{"Date", "Time", "Vab", "Vbc", "Vca", "Ia", "Ib", "Ic", "MW", "MVAR", "PF"}
	for _, header := range headers {
		pdf.CellFormat(18, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 8)

	for i, data := range dailyData {
		if i == 22 || i == 47 {
			pdf.SetFont("Arial", "B", 10)
			headers := []string{"Date", "Time", "Vac", "Vbc", "Vca", "Ia", "Ib", "Ic", "MW", "MVAR", "PF"}
			for _, header := range headers {
				pdf.CellFormat(18, 10, header, "1", 0, "C", false, 0, "")
			}
			pdf.SetFont("Arial", "", 8)
			pdf.Ln(-1)
		}
		pdf.CellFormat(18, 10, data.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 10, data.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 10, data.Vab, "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 10, data.Vbc, "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 10, data.Vca, "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 10, data.CurrentPhaseA, "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 10, data.CurrentPhaseB, "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 10, data.CurrentPhaseC, "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 10, data.ActivePower, "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 10, data.ReactivePower, "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 10, data.PowerFactor, "1", 0, "C", false, 0, "")
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
	pdf.SetFont("Arial", "B", 5)
	pdf.CellFormat(285, 5, "Monthly Load Report", "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(285, 5, substationName, "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(285, 9, bName, "LTRB", 0, "C", false, 0, "")
	// pdf.SetY(35)

	pdf.Ln(-1)
	pdf.CellFormat(95, 9, "DAY TIME PEAK", "LTRB", 0, "C", false, 0, "")
	pdf.CellFormat(95, 9, "NIGHT TIME PEAK", "LTRB", 0, "C", false, 0, "")
	pdf.CellFormat(95, 9, "DAY&NIGHT LIGHT LOAD", "LTRB", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(95, 9, "08:00-15:30", "LTRB", 0, "C", false, 0, "")
	pdf.CellFormat(95, 9, "00:00-07:30,16:00-23:30", "LTRB", 0, "C", false, 0, "")
	pdf.CellFormat(95, 9, "00:00-23:30", "LTRB", 0, "C", false, 0, "")
	pdf.Ln(-1)

	// Header
	headers := []string{"Date", "Time", "Vab (kv)", "Vbc (kv)", "Vca (kv)", "Ia (A)", "Ib (A)", "Ic (A)", "P (MW)", "Q (MVAR)", "PF (%)"}
	for i, header := range headers {
		if i == 0 {
			pdf.CellFormat(15, 7, header, "1", 0, "C", false, 0, "")
		} else {
			pdf.CellFormat(8, 7, header, "1", 0, "C", false, 0, "")
		}
	}
	for i, header := range headers {
		if i == 0 {
			pdf.CellFormat(15, 7, header, "1", 0, "C", false, 0, "")
		} else {
			pdf.CellFormat(8, 7, header, "1", 0, "C", false, 0, "")
		}
	}
	for i, header := range headers {
		if i == 0 {
			pdf.CellFormat(15, 7, header, "1", 0, "C", false, 0, "")
		} else {
			pdf.CellFormat(8, 7, header, "1", 0, "C", false, 0, "")
		}
	}

	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 5)
	for i, data := range items {
		pdf.CellFormat(15, 7, data.PeakDay.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakDay.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakDay.Vab, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakDay.Vbc, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakDay.Vca, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakDay.Ia, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakDay.Ib, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakDay.Ic, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakDay.Mw, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakDay.Mvar, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakDay.P, "1", 0, "C", false, 0, "")

		pdf.CellFormat(15, 7, data.PeakNight.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakNight.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakNight.Vab, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakNight.Vbc, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakNight.Vca, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakNight.Ia, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakNight.Ib, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakNight.Ic, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakNight.Mw, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakNight.Mvar, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.PeakNight.P, "1", 0, "C", false, 0, "")

		pdf.CellFormat(15, 7, data.All.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.All.Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.All.Vab, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.All.Vbc, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.All.Vca, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.All.Ia, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.All.Ib, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.All.Ic, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.All.Mw, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.All.Mvar, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 7, data.All.P, "1", 0, "C", false, 0, "")

		if i == 18 {
			pdf.Ln(-1)
			for i, header := range headers {
				if i == 0 {
					pdf.CellFormat(15, 7, header, "1", 0, "C", false, 0, "")
				} else {
					pdf.CellFormat(8, 7, header, "1", 0, "C", false, 0, "")
				}
			}
			for i, header := range headers {
				if i == 0 {
					pdf.CellFormat(15, 7, header, "1", 0, "C", false, 0, "")
				} else {
					pdf.CellFormat(8, 7, header, "1", 0, "C", false, 0, "")
				}
			}
			for i, header := range headers {
				if i == 0 {
					pdf.CellFormat(15, 7, header, "1", 0, "C", false, 0, "")
				} else {
					pdf.CellFormat(8, 7, header, "1", 0, "C", false, 0, "")
				}
			}
		}
		pdf.Ln(-1)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)

	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func ExportPdfYearly(peak []dto.DataTmpsYear, light []dto.DataTmpsYear, sName string, bName string, exportHeader string, year int) (*bytes.Buffer, error) {

	substationName := exportHeader
	bayName := bName
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	imagePath := "static/css/icons/image.png"
	pdf.Image(imagePath, 10, 15, 20, 0, false, "", 0, "")
	// Set font
	pdf.SetFont("Arial", "B", 7)
	pdf.CellFormat(192, 10, "Yearly Load Report", "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(192, 10, substationName, "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(192, 10, fmt.Sprintf("%d", year), "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(192, 10, bayName, "1", 0, "C", false, 0, "")

	pdf.Ln(-1)
	pdf.CellFormat(96, 10, "Peak Load", "1", 0, "C", false, 0, "")
	pdf.CellFormat(96, 10, "Light Load", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	// Header
	pdf.SetFont("Arial", "B", 5)
	headers := []string{"Month", "Date", "Time", "Vab (kV)", "Vbc (kV)", "Vca (kV)", "Ia (A)", "Ib (A)", "Ic (A)", "P (MW)", "Q (MVAR)", "PF (%)"}
	for i := range headers {
		pdf.CellFormat(8, 10, headers[i], "1", 0, "C", false, 0, "")
	}
	for i := range headers {
		pdf.CellFormat(8, 10, headers[i], "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 5)

	for i := range peak {
		pdf.CellFormat(8, 10, peak[i].Month, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, peak[i].Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, peak[i].Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, peak[i].Vab, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, peak[i].Vbc, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, peak[i].Vca, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, peak[i].CurrentPhaseA, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, peak[i].CurrentPhaseB, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, peak[i].CurrentPhaseC, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, peak[i].ActivePower, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, peak[i].ReactivePower, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, peak[i].PowerFactor, "1", 0, "C", false, 0, "")

		pdf.CellFormat(8, 10, light[i].Month, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, light[i].Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, light[i].Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, light[i].Vab, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, light[i].Vbc, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, light[i].Vca, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, light[i].CurrentPhaseA, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, light[i].CurrentPhaseB, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, light[i].CurrentPhaseC, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, light[i].ActivePower, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, light[i].ReactivePower, "1", 0, "C", false, 0, "")
		pdf.CellFormat(8, 10, peak[i].PowerFactor, "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)

	if err != nil {
		return nil, err
	}
	return &buf, nil
}
