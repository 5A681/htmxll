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
	ExportExcelDaily(dailyData []dto.DataTmps, fileName string) error
	ExportExcelMonthly(day []dto.DataTmps, night []dto.DataTmps, all []dto.DataTmps, fileName string, subS string, bay string) error
	ExportExcelYearly(peak []dto.DataTmpsYear, light []dto.DataTmpsYear, fileName string, subS string, bay string) error
}
type exportExcel struct {
	excel *excelize.File
}

func NewExportExcel(excel *excelize.File) ExportExcel {
	return exportExcel{excel}
}

func (e exportExcel) CreateSheetYearly(excel *excelize.File, datas []dto.DataTmpsYear, fileName string, sheetName string, title string, subStation string, bay string, timeTitle string) error {
	index, _ := excel.NewSheet(sheetName)

	// Set table headers
	headers := []string{"Month", "Date", "Time", "Vab (kV)", "Vbc (kV)", "Vca (kV)", "Ia (A)", "Ib (A)", "Ic (A)", "P (PW)", "Q (MVAR)", "PF (%)"}
	for i, header := range headers {
		cell := string(rune('A'+i)) + "9" // A1, B1, C1, etc.
		excel.SetCellValue(sheetName, cell, header)
	}

	centerRowStyleId, err := excel.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		return err
	}

	if err := excel.SetRowStyle(sheetName, 4, 5, centerRowStyleId); err != nil {
		return err
	}

	Title := title
	SubSstationName := subStation
	bayName := bay
	if err := excel.SetCellValue(sheetName, "A4", Title); err != nil {
		return err
	}
	if err := excel.SetCellValue(sheetName, "A5", SubSstationName); err != nil {
		return err
	}

	if err := excel.SetCellValue(sheetName, "A6", bayName); err != nil {
		return err
	}
	if err := excel.SetCellValue(sheetName, "A7", timeTitle); err != nil {
		return err
	}

	for row, data := range datas {
		cell := "A" + fmt.Sprintf("%d", 9+row)
		excel.SetCellValue(sheetName, cell, data.Month)
		cell = "B" + fmt.Sprintf("%d", 9+row)
		excel.SetCellValue(sheetName, cell, data.Date)
		cell = "C" + fmt.Sprintf("%d", 9+row)
		excel.SetCellValue(sheetName, cell, data.Time)
		cell = "D" + fmt.Sprintf("%d", 9+row)
		excel.SetCellValue(sheetName, cell, "")
		cell = "E" + fmt.Sprintf("%d", 9+row)
		excel.SetCellValue(sheetName, cell, "")
		cell = "F" + fmt.Sprintf("%d", 9+row)
		excel.SetCellValue(sheetName, cell, "")
		cell = "G" + fmt.Sprintf("%d", 9+row)
		excel.SetCellValue(sheetName, cell, data.CurrentPhaseA)
		cell = "H" + fmt.Sprintf("%d", 9+row)
		excel.SetCellValue(sheetName, cell, data.CurrentPhaseB)
		cell = "I" + fmt.Sprintf("%d", 9+row)
		excel.SetCellValue(sheetName, cell, data.CurrentPhaseC)
		cell = "J" + fmt.Sprintf("%d", 9+row)
		excel.SetCellValue(sheetName, cell, data.ActivePower)
		cell = "K" + fmt.Sprintf("%d", 9+row)
		excel.SetCellValue(sheetName, cell, data.ReactivePower)
		cell = "L" + fmt.Sprintf("%d", 9+row)
		excel.SetCellValue(sheetName, cell, data.PowerFactor)
	}

	if err := excel.MergeCell(sheetName, "A4", "L4"); err != nil {
		return err
	}
	if err := excel.MergeCell(sheetName, "A5", "L5"); err != nil {
		return err
	}
	if err := excel.MergeCell(sheetName, "A6", "L6"); err != nil {
		return err
	}
	if err := excel.MergeCell(sheetName, "A7", "L7"); err != nil {
		return err
	}

	// Define the table range
	tableRange := "A8:L56" // Includes headers and data

	// Create a table with the defined range
	disable := true

	if err := excel.AddTable(sheetName, &excelize.Table{
		Range:             tableRange,
		Name:              "Table1",
		StyleName:         "TableStyleMedium9",
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
	if err := excel.AddPicture(sheetName, "A3", imagePath, &excelize.GraphicOptions{
		AutoFit: false,
		OffsetX: 10,
		OffsetY: 10,
		ScaleX:  0.3,
		ScaleY:  0.3,
	}); err != nil {
		return err
	}

	if err := excel.SetColWidth(sheetName, "A", "L", 15); err != nil {
		return err
	}

	stypeId, err := excel.NewStyle(&excelize.Style{
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
	err = excel.SetCellStyle(sheetName, "A6", "L56", stypeId)
	if err != nil {
		return err
	}

	// Set active sheet and save the file
	excel.SetActiveSheet(index)

	err = excel.DeleteTable("Table1")
	if err != nil {
		return err
	}
	return nil
}

func (e exportExcel) CreateSheetMonthly(excel *excelize.File, datas []dto.DataTmps, fileName string, sheetName string, title string, subStation string, bay string, timeTitle string, timeRange string) error {
	index, _ := excel.NewSheet(sheetName)

	// Set table headers
	headers := []string{"Date", "Time", "Vab (kV)", "Vbc (kV)", "Vca (kV)", "Ia (A)", "Ib (A)", "Ic (A)", "P (PW)", "Q (MVAR)", "PF (%)"}
	for i, header := range headers {
		cell := string(rune('A'+i)) + "9" // A1, B1, C1, etc.
		excel.SetCellValue(sheetName, cell, header)
	}

	centerRowStyleId, err := excel.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		return err
	}

	if err := excel.SetRowStyle(sheetName, 4, 5, centerRowStyleId); err != nil {
		return err
	}

	Title := title
	SubSstationName := subStation
	bayName := bay
	if err := excel.SetCellValue(sheetName, "A4", Title); err != nil {
		return err
	}
	if err := excel.SetCellValue(sheetName, "A5", SubSstationName); err != nil {
		return err
	}

	if err := excel.SetCellValue(sheetName, "A6", bayName); err != nil {
		return err
	}
	if err := excel.SetCellValue(sheetName, "A7", timeTitle); err != nil {
		return err
	}
	if err := excel.SetCellValue(sheetName, "A8", timeRange); err != nil {
		return err
	}

	for row, data := range datas {
		cell := "A" + fmt.Sprintf("%d", 10+row)
		excel.SetCellValue(sheetName, cell, data.Date)
		cell = "B" + fmt.Sprintf("%d", 10+row)
		excel.SetCellValue(sheetName, cell, data.Time)
		cell = "C" + fmt.Sprintf("%d", 10+row)
		excel.SetCellValue(sheetName, cell, "")
		cell = "D" + fmt.Sprintf("%d", 10+row)
		excel.SetCellValue(sheetName, cell, "")
		cell = "E" + fmt.Sprintf("%d", 10+row)
		excel.SetCellValue(sheetName, cell, "")
		cell = "F" + fmt.Sprintf("%d", 10+row)
		excel.SetCellValue(sheetName, cell, data.CurrentPhaseA)
		cell = "G" + fmt.Sprintf("%d", 10+row)
		excel.SetCellValue(sheetName, cell, data.CurrentPhaseB)
		cell = "H" + fmt.Sprintf("%d", 10+row)
		excel.SetCellValue(sheetName, cell, data.CurrentPhaseC)
		cell = "I" + fmt.Sprintf("%d", 10+row)
		excel.SetCellValue(sheetName, cell, data.ActivePower)
		cell = "J" + fmt.Sprintf("%d", 10+row)
		excel.SetCellValue(sheetName, cell, data.ReactivePower)
		cell = "K" + fmt.Sprintf("%d", 10+row)
		excel.SetCellValue(sheetName, cell, data.PowerFactor)
	}

	if err := excel.MergeCell(sheetName, "A4", "K4"); err != nil {
		return err
	}
	if err := excel.MergeCell(sheetName, "A5", "K5"); err != nil {
		return err
	}
	if err := excel.MergeCell(sheetName, "A6", "K6"); err != nil {
		return err
	}
	if err := excel.MergeCell(sheetName, "A7", "K7"); err != nil {
		return err
	}
	if err := excel.MergeCell(sheetName, "A8", "K8"); err != nil {
		return err
	}

	// Define the table range
	tableRange := "A9:K57" // Includes headers and data

	// Create a table with the defined range
	disable := true

	if err := excel.AddTable(sheetName, &excelize.Table{
		Range:             tableRange,
		Name:              "Table1",
		StyleName:         "TableStyleMedium9",
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
	if err := excel.AddPicture(sheetName, "A3", imagePath, &excelize.GraphicOptions{
		AutoFit: false,
		OffsetX: 10,
		OffsetY: 10,
		ScaleX:  0.3,
		ScaleY:  0.3,
	}); err != nil {
		return err
	}

	if err := excel.SetColWidth(sheetName, "A", "K", 15); err != nil {
		return err
	}

	stypeId, err := excel.NewStyle(&excelize.Style{
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
	err = excel.SetCellStyle(sheetName, "A6", "K57", stypeId)
	if err != nil {
		return err
	}

	// Set active sheet and save the file
	excel.SetActiveSheet(index)

	err = excel.DeleteTable("Table1")
	if err != nil {
		return err
	}
	return nil
}

func (e exportExcel) ExportExcelYearly(peak []dto.DataTmpsYear, light []dto.DataTmpsYear, fileName string, subS string, bay string) error {
	excel := e.excel
	err := e.CreateSheetYearly(excel, peak, fileName, "Peak Load", "Yearly Load Report", subS, bay, "Peak Load")
	if err != nil {
		return err
	}
	err = e.CreateSheetYearly(excel, light, fileName, "Light Load", "Yearly Load Report", subS, bay, "Light Load")
	if err != nil {
		return err
	}
	if err = excel.SaveAs(fileName); err != nil {
		return err
	}

	return nil
}

func (e exportExcel) ExportExcelMonthly(day []dto.DataTmps, night []dto.DataTmps, all []dto.DataTmps, fileName string, subS string, bay string) error {
	excel := e.excel
	err := e.CreateSheetMonthly(excel, day, fileName, "Day Time Peak", "Monthly Load Report", subS, bay, "Day Time Peak", "08:00-15:30")
	if err != nil {
		return err
	}
	err = e.CreateSheetMonthly(excel, night, fileName, "Night Time Peak", "Monthly Load Report", subS, bay, "Night Time Peak", "00:00-07:30, 16:00-23:30")
	if err != nil {
		return err
	}
	err = e.CreateSheetMonthly(excel, all, fileName, "Day & Night Light Load", "Monthly Load Report", subS, bay, "Day & Night Light Load Peak", "00:00-23:30")
	if err != nil {
		return err
	}
	if err = excel.SaveAs(fileName); err != nil {
		return err
	}

	return nil
}

func (e exportExcel) ExportExcelDaily(dailyData []dto.DataTmps, fileName string) error {

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
	SubSstationName := "Substation Name"
	bayName := "Bay Name"
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
		e.excel.SetCellValue(sheetName, cell, "")
		cell = "D" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, "")
		cell = "E" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, "")
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
		StyleName:         "TableStyleMedium9",
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
