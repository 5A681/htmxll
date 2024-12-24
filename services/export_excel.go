package services

import (
	"fmt"
	"htmxll/dto"
	_ "image/png"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

type ExportExcel interface {
	ExportExcelDaily(dailyData []dto.DataTmps, fileName string, bayName string, exportHeader string) error
	ExportExcelMonthly(items []dto.MonthlyRowData, fileName string, subS string, bay string, exportHeader string) error
	ExportExcelYearly(peak []dto.DataTmpsYear, light []dto.DataTmpsYear, fileName string, subS string, bay string, year int) error
}
type exportExcel struct {
	excel *excelize.File
}

func NewExportExcel(excel *excelize.File) ExportExcel {
	return exportExcel{excel}
}
func getExcelColumnName(n int) string {
	columnName := ""
	for n > 0 {
		n-- // Adjust to zero-based index
		columnName = string('A'+(n%26)) + columnName
		n /= 26
	}
	return columnName
}
func (e exportExcel) ExportExcelYearly(peak []dto.DataTmpsYear, light []dto.DataTmpsYear, fileName string, subS string, bay string, year int) error {

	sheetName := "Sheet1"

	index, _ := e.excel.NewSheet(sheetName)

	tableRange := "A9:X55" // Includes headers and data

	// Create a table with the defined range
	disable := false

	if err := e.excel.AddTable(sheetName, &excelize.Table{
		Range:             tableRange,
		Name:              "Table1",
		StyleName:         "TableStyleMedium1",
		ShowFirstColumn:   false,
		ShowLastColumn:    false,
		ShowRowStripes:    &disable,
		ShowColumnStripes: false,
	}); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Set table headers
	headers := []string{"Month", "Date", "Time", "Vab (kV)", "Vbc (kV)", "Vca (kV)", "Ia (A)", "Ib (A)", "Ic (A)", "P (PW)", "Q (MVAR)", "PF (%)", "Month", "Date", "Time", "Vab (kV)", "Vbc (kV)", "Vca (kV)", "Ia (A)", "Ib (A)", "Ic (A)", "P (PW)", "Q (MVAR)", "PF (%)"}

	row := 9
	for i, header := range headers {
		column := getExcelColumnName(i + 1) // Convert index to column name (1-based)
		cell := fmt.Sprintf("%s%d", column, row)
		if err := e.excel.SetCellValue(sheetName, cell, header); err != nil {
			log.Fatalf("Failed to set cell value for %s: %v", cell, err)
		}
	}

	centerRowStyleId, err := e.excel.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		return err
	}

	if err := e.excel.SetRowStyle(sheetName, 4, 5, centerRowStyleId); err != nil {
		return err
	}

	Title := "Daily Load Report"
	SubSstationName := subS
	if err := e.excel.SetCellValue(sheetName, "A4", Title); err != nil {
		return err
	}
	if err := e.excel.SetCellValue(sheetName, "A5", SubSstationName); err != nil {
		return err
	}

	if err := e.excel.SetCellValue(sheetName, "A6", fmt.Sprintf("%d", year)); err != nil {
		return err
	}
	if err := e.excel.SetCellValue(sheetName, "A7", bay); err != nil {
		return err
	}
	if err := e.excel.SetCellValue(sheetName, "A8", "Peak Load"); err != nil {
		return err
	}
	if err := e.excel.SetCellValue(sheetName, "M8", "Loght Load"); err != nil {
		return err
	}

	for row := range peak {
		cell := "A" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, peak[row].Month)
		cell = "B" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, peak[row].Date)
		cell = "C" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, peak[row].Time)
		cell = "D" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, peak[row].Vab)
		cell = "E" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, peak[row].Vbc)
		cell = "F" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, peak[row].Vca)
		cell = "G" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, peak[row].CurrentPhaseA)
		cell = "H" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, peak[row].CurrentPhaseB)
		cell = "I" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, peak[row].CurrentPhaseC)
		cell = "J" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, peak[row].ActivePower)
		cell = "K" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, peak[row].ReactivePower)
		cell = "L" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, peak[row].PowerFactor)

		//agian
		cell = "M" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, light[row].Month)
		cell = "N" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, light[row].Date)
		cell = "O" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, light[row].Time)
		cell = "P" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, light[row].Vab)
		cell = "Q" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, light[row].Vbc)
		cell = "R" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, light[row].Vca)
		cell = "S" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, light[row].CurrentPhaseA)
		cell = "T" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, light[row].CurrentPhaseB)
		cell = "U" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, light[row].CurrentPhaseC)
		cell = "V" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, light[row].ActivePower)
		cell = "W" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, light[row].ReactivePower)
		cell = "X" + fmt.Sprintf("%d", 10+row)
		e.excel.SetCellValue(sheetName, cell, light[row].PowerFactor)

	}

	if err := e.excel.MergeCell(sheetName, "A4", "X4"); err != nil {
		return err
	}
	if err := e.excel.MergeCell(sheetName, "A5", "X5"); err != nil {
		return err
	}
	if err := e.excel.MergeCell(sheetName, "A6", "X6"); err != nil {
		return err
	}
	if err := e.excel.MergeCell(sheetName, "A7", "X7"); err != nil {
		return err
	}
	if err := e.excel.MergeCell(sheetName, "A8", "L8"); err != nil {
		return err
	}
	if err := e.excel.MergeCell(sheetName, "M8", "X8"); err != nil {
		return err
	}

	// Define the table range

	// Set bold borders for the table
	// boldBorderStyle := excelize.Border{
	// 	Type:  "top",
	// 	Color: "000000",
	// 	Style: 2, // Bold border style
	// }
	imagePath := "static/css/icons/image.png"

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		log.Fatalf("Image file does not exist: %v", err)
		return err
	}
	if err := e.excel.AddPicture(sheetName, "A3", imagePath, &excelize.GraphicOptions{
		AutoFit: false,
		OffsetX: 10,
		OffsetY: 10,
		ScaleX:  0.3,
		ScaleY:  0.3,
	}); err != nil {
		return err
	}

	if err := e.excel.SetColWidth(sheetName, "A", "X", 15); err != nil {
		return err
	}

	stypeId, err := e.excel.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 2},
			{Type: "top", Color: "000000", Style: 2},
			{Type: "bottom", Color: "000000", Style: 2},
			{Type: "right", Color: "000000", Style: 2},
		},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		return err
	}
	// Apply bold borders to all sides
	err = e.excel.SetCellStyle(sheetName, "A6", "X55", stypeId)
	if err != nil {
		return err
	}

	// Set active sheet and save the file
	e.excel.SetActiveSheet(index)
	if err = e.excel.SaveAs(fileName); err != nil {
		return err
	}

	err = e.excel.DeleteTable("Table1")
	if err != nil {
		return err
	}
	return nil

}

func (e exportExcel) ExportExcelMonthly(items []dto.MonthlyRowData, fileName string, subS string, bay string, exportHeader string) error {
	f := excelize.NewFile()
	defer f.Close()
	sheetName := "Sheet1"
	tableRange := "A9:AG40" // Includes headers and data

	//Create a table with the defined range
	disable := true

	if err := f.AddTable(sheetName, &excelize.Table{
		Range:             tableRange,
		Name:              "Table1",
		StyleName:         "TableStyleMedium1",
		ShowFirstColumn:   false,
		ShowLastColumn:    false,
		ShowRowStripes:    &disable,
		ShowColumnStripes: false,
	}); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	index, _ := f.NewSheet(sheetName)
	if err := f.MergeCell(sheetName, "A7", "K7"); err != nil {
		return err
	}
	if err := f.MergeCell(sheetName, "L7", "V7"); err != nil {
		return err
	}
	if err := f.MergeCell(sheetName, "W7", "AG7"); err != nil {
		return err
	}
	if err := f.MergeCell(sheetName, "A8", "K8"); err != nil {
		return err
	}
	if err := f.MergeCell(sheetName, "L8", "V8"); err != nil {
		return err
	}
	if err := f.MergeCell(sheetName, "W8", "AG8"); err != nil {
		return err
	}

	if err := f.MergeCell(sheetName, "A4", "AG4"); err != nil {
		return err
	}
	if err := f.MergeCell(sheetName, "A5", "AG5"); err != nil {
		return err
	}
	if err := f.MergeCell(sheetName, "A6", "AG6"); err != nil {
		return err
	}

	//Define the table range

	//Set table headers
	headers := []string{"08.00-15.30", "00:00-07:30,16.00-23.30", "00.00-23.30"}
	f.SetCellValue(sheetName, "A8", headers[0])
	f.SetCellValue(sheetName, "L8", headers[1])
	f.SetCellValue(sheetName, "W8", headers[2])

	f.SetCellValue(sheetName, "A7", "Day Time Peak")
	f.SetCellValue(sheetName, "L7", "Night Time Peak")
	f.SetCellValue(sheetName, "W7", "Day & Night Light Load")

	subheaders := []string{"Date", "Time", "Vab (kv)", "Vbc (kv)", "Vca (kv)", "Ia (A)", "Ib (A)", "Ic (A)", " P (MW)", "Q (MVAR)", "PF (%)"}
	for j, sub := range subheaders {
		cell := string(rune('A'+j)) + "9" // A1, B1, C1, etc.
		f.SetCellValue(sheetName, cell, sub)
	}
	for j, sub := range subheaders {
		cell := string(rune('L'+j)) + "9" // A1, B1, C1, etc.
		f.SetCellValue(sheetName, cell, sub)
	}
	for j, sub := range subheaders {
		cell := string(rune('W'+j)) + "9" // A1, B1, C1, etc.
		if j >= 4 {
			cell = "A" + string(rune('A'+j-4)) + "9"
		}
		f.SetCellValue(sheetName, cell, sub)
	}

	for row, data := range items {

		cell := "A" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakDay.Date)
		cell = "B" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakDay.Time)
		cell = "C" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakDay.Vab)
		cell = "D" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakDay.Vbc)
		cell = "E" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakDay.Vca)
		cell = "F" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakDay.Ia)
		cell = "G" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakDay.Ib)
		cell = "H" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakDay.Ic)
		cell = "I" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakDay.Mw)
		cell = "J" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakDay.Mvar)
		cell = "K" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakDay.P)

		cell = "L" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakNight.Date)
		cell = "M" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakNight.Time)
		cell = "N" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakNight.Vab)
		cell = "O" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakNight.Vbc)
		cell = "P" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakNight.Vca)
		cell = "Q" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakNight.Ia)
		cell = "R" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakNight.Ib)
		cell = "S" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakNight.Ic)
		cell = "T" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakNight.Mw)
		cell = "U" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakNight.Mvar)
		cell = "V" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.PeakNight.P)

		cell = "W" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.All.Date)
		cell = "X" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.All.Time)
		cell = "Y" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.All.Vab)
		cell = "Z" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.All.Vbc)
		cell = "AA" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.All.Vca)
		cell = "AB" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.All.Ia)
		cell = "AC" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.All.Ib)
		cell = "AD" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.All.Ic)
		cell = "AE" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.All.Mw)
		cell = "AF" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.All.Mvar)
		cell = "AG" + fmt.Sprintf("%d", 10+row)
		f.SetCellValue(sheetName, cell, data.All.P)

	}

	centerRowStyleId, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		return err
	}

	if err := f.SetRowStyle(sheetName, 4, 5, centerRowStyleId); err != nil {
		return err
	}

	Title := "Monthly Load Report"
	SubSstationName := exportHeader

	if err := f.SetCellValue(sheetName, "A4", Title); err != nil {
		return err
	}

	if err := f.SetCellValue(sheetName, "A5", SubSstationName); err != nil {
		return err
	}
	if err := f.SetCellValue(sheetName, "A6", bay); err != nil {
		return err
	}

	// Set bold borders for the table
	// boldBorderStyle := excelize.Border{
	// 	Type:  "top",
	// 	Color: "000000",
	// 	Style: 2, // Bold border style
	// }
	imagePath := "static/css/icons/image.png"

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		log.Fatalf("Image file does not exist: %v", err)
		return err
	}
	if err := f.AddPicture(sheetName, "A3", imagePath, &excelize.GraphicOptions{
		AutoFit: false,
		OffsetX: 10,
		OffsetY: 10,
		ScaleX:  0.3,
		ScaleY:  0.3,
	}); err != nil {
		return err
	}

	// if err := e.excel.SetColWidth(sheetName, "A", "AB", 15); err != nil {
	// 	return err
	// }

	stypeId, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 2},
			{Type: "top", Color: "000000", Style: 2},
			{Type: "bottom", Color: "000000", Style: 2},
			{Type: "right", Color: "000000", Style: 2},
		},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		return err
	}
	// Apply bold borders to all sides
	err = f.SetCellStyle(sheetName, "A6", "AG40", stypeId)
	if err != nil {
		return err
	}

	// Set active sheet and save the file
	f.SetActiveSheet(index)
	if err = f.SaveAs(fileName); err != nil {
		return err
	}

	// err = e.excel.DeleteTable("Table1")
	// if err != nil {
	// 	return err
	// }

	return nil

}

func (e exportExcel) ExportExcelDaily(dailyData []dto.DataTmps, fileName string, bayName string, exportHeader string) error {

	sheetName := "Sheet1"

	index, _ := e.excel.NewSheet(sheetName)

	// Set table headers
	headers := []string{"Date", "Time", "Vab (kV)", "Vbc (kV)", "Vca (kV)", "Ia (A)", "Ib (A)", "Ic (A)", "P (PW)", "Q (MVAR)", "PF (%)"}
	for i, header := range headers {
		cell := string(rune('A'+i)) + "7" // A1, B1, C1, etc.
		e.excel.SetCellValue(sheetName, cell, header)
	}

	centerRowStyleId, err := e.excel.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		return err
	}

	if err := e.excel.SetRowStyle(sheetName, 4, 5, centerRowStyleId); err != nil {
		return err
	}

	Title := "Daily Load Report"
	SubSstationName := exportHeader
	if err := e.excel.SetCellValue(sheetName, "A4", Title); err != nil {
		return err
	}
	if err := e.excel.SetCellValue(sheetName, "A5", SubSstationName); err != nil {
		return err
	}

	if err := e.excel.SetCellValue(sheetName, "A6", bayName); err != nil {
		return err
	}

	// Set table data
	// data := [][]interface{}{
	// 	{1, "John Doe", 25, 1},
	// 	{2, "Jane Smith", 30, 2},
	// 	{3, "Sam Brown", 22, 3},
	// }

	// for i, row := range data {
	// 	for j, value := range row {
	// 		cell := string(rune('A'+j)) + fmt.Sprintf("%d", 8+i)
	// 		fmt.Println(cell, value)
	// 		e.excel.SetCellValue(sheetName, cell, value)
	// 	}
	// }
	for row, data := range dailyData {
		cell := "A" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.Date)
		cell = "B" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.Time)
		cell = "C" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.Vab)
		cell = "D" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.Vbc)
		cell = "E" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.Vca)
		cell = "F" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.CurrentPhaseA)
		cell = "G" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.CurrentPhaseB)
		cell = "H" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.CurrentPhaseC)
		cell = "I" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.ActivePower)
		cell = "J" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.ReactivePower)
		cell = "K" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.PowerFactor)
	}

	if err := e.excel.MergeCell(sheetName, "A4", "K4"); err != nil {
		return err
	}
	if err := e.excel.MergeCell(sheetName, "A5", "K5"); err != nil {
		return err
	}
	if err := e.excel.MergeCell(sheetName, "A6", "K6"); err != nil {
		return err
	}

	// Define the table range
	tableRange := "A7:K55" // Includes headers and data

	// Create a table with the defined range
	disable := true

	if err := e.excel.AddTable(sheetName, &excelize.Table{
		Range:             tableRange,
		Name:              "Table1",
		StyleName:         "TableStyleMedium1",
		ShowFirstColumn:   false,
		ShowLastColumn:    false,
		ShowRowStripes:    &disable,
		ShowColumnStripes: false,
	}); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Set bold borders for the table
	// boldBorderStyle := excelize.Border{
	// 	Type:  "top",
	// 	Color: "000000",
	// 	Style: 2, // Bold border style
	// }
	imagePath := "static/css/icons/image.png"

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		log.Fatalf("Image file does not exist: %v", err)
		return err
	}
	if err := e.excel.AddPicture(sheetName, "A3", imagePath, &excelize.GraphicOptions{
		AutoFit: false,
		OffsetX: 10,
		OffsetY: 10,
		ScaleX:  0.3,
		ScaleY:  0.3,
	}); err != nil {
		return err
	}

	if err := e.excel.SetColWidth(sheetName, "A", "K", 15); err != nil {
		return err
	}

	stypeId, err := e.excel.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 2},
			{Type: "top", Color: "000000", Style: 2},
			{Type: "bottom", Color: "000000", Style: 2},
			{Type: "right", Color: "000000", Style: 2},
		},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		return err
	}
	// Apply bold borders to all sides
	err = e.excel.SetCellStyle(sheetName, "A6", "K55", stypeId)
	if err != nil {
		return err
	}

	// Set active sheet and save the file
	e.excel.SetActiveSheet(index)
	if err = e.excel.SaveAs(fileName); err != nil {
		return err
	}

	err = e.excel.DeleteTable("Table1")
	if err != nil {
		return err
	}
	return nil
}
