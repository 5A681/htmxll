package filedata

import (
	"htmxll/entity"
	"htmxll/repository"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/extrame/xls"
)

func ReadFileXls(filePath string, sheet int, dataTempRepo repository.Repository) {
	xlsFile, err := xls.Open(filePath, "utf-8")
	if err != nil {
		log.Printf("Failed to open .xls file: %v", err)
		return
	}
	if xlsFile == nil {
		return
	}
	_ = xlsFile
	ws := xlsFile.GetSheet(sheet)
	if ws == nil {
		return
	}
	maxRow := int(ws.MaxRow)
	if maxRow <= 0 {
		return
	}
	maxCol := ws.Row(0).LastCol()
	stationAndBay := ws.Row(1).Col(2)

	splitName := strings.Split(stationAndBay, ".")
	if len(splitName) < 2 {
		return
	}
	log.Printf("name = %s, %s, %s\n", splitName[0], splitName[1], splitName[2])
	if splitName[0] != "" {
		_, err := dataTempRepo.GetSubStationByName(splitName[0])
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				err = dataTempRepo.CreateSubStation(&entity.SubStation{Name: splitName[0]})
				if err != nil {
					log.Println(err)
				}
			} else {
				log.Println(err)
			}
		}

	}
	subId := 0
	if ws.Name != "" {
		sub, err := dataTempRepo.GetSubStationByName(splitName[0])
		if err != nil {
			log.Println(err)
			return
		}
		subId = sub.Id
		_, err = dataTempRepo.GetBayByNameAndSubStationId(sub.Id, ws.Name)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				log.Println("bay name = ", ws.Name)
				err = dataTempRepo.CreateBay(&entity.Bay{
					Name:         ws.Name,
					SubStationId: sub.Id,
					SheetNumber:  sheet,
				})
				if err != nil {
					log.Println("could not create bay", err, subId*sheet, subId, sheet)
				}
			} else {
				log.Println(err)
			}
		}

	}
	if subId == 0 {
		return
	}
	bay, err := dataTempRepo.GetBayByNameAndSubStationId(subId, ws.Name)
	if err != nil {
		log.Println("error get bay", err)
		return
	}
	for r := 5; r < maxRow; r++ {

		if ws.Row(r).Col(0) != "" {

			tempData := entity.DataTmps{
				BayId: bay.Id,
			}
			for c := 0; c < maxCol; c += 2 {

				if c == 0 {
					//fmt.Printf("%v \t", ReadDateTimeColumn(ws.Row(r).Col(c)))
					dateTime := ReadDateTimeColumn(ws.Row(r).Col(c))
					if dateTime != nil {
						tempData.DataDatetime = *dateTime
					}
				}
				columnName := ws.Row(2).Col(c)
				if columnName != "" {
					floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
					if err != nil {
						MapToInsert(0, columnName, &tempData)
					} else {
						data := float32(floatData)
						MapToInsert(data, columnName, &tempData)
					}
				}

			}
			tempData.CreatedAt = time.Now()
			err := dataTempRepo.CreateDataTmep(tempData)
			if err != nil {
				log.Println("could not insert temp data", err.Error())
			}
		} else {
			return
		}

	}
}

func MapToInsert(data float32, field string, model *entity.DataTmps) {
	switch field {
	case "BUS VOLTAGE A-B":
		model.VoltageAB = data
	case "BUS VOLTAGE B-C":
		model.VoltageBC = data
	case "BUS VOLTAGE C-A":
		model.VoltageCA = data
	case "VOLTAGE A-B":
		model.VoltageAB = data
	case "VOLTAGE B-C":
		model.VoltageBC = data
	case "VOLTAGE C-A":
		model.VoltageCA = data
	case "CURRENT PHASE A":
		model.CurrentPhaseA = data
	case "CURRENT PHASE B":
		model.CurrentPhaseB = data
	case "CURRENT PHASE C":
		model.CurrentPhaseC = data
	case "ACTIVE POWER P":
		model.ActivePower = data
	case "ACTIVE POWER":
		model.ActivePower = data
	case "REACTIVE POWER Q":
		model.ReactivePower = data
	case "REACTIVE POWER":
		model.ReactivePower = data
	case "POWER FACTOR PF":
		model.PowerFactor = data
	case "POWER FACTOR":
		model.PowerFactor = data
	default:
		return
	}
}

func ReadDateTimeColumn(datetime string) *time.Time {
	var realTime *time.Time
	t, err := strconv.ParseFloat(datetime, 64)
	if err != nil {
		convertTime, err := time.Parse(time.RFC3339, datetime)
		if err != nil {
			return nil
		} else {
			if convertTime.Second() != 0 {
				convertTime = convertTime.Add(time.Second * (time.Duration(60 - convertTime.Second())))
			}
			realTime = &convertTime
		}

	} else {
		//realTime =
		convertTime := ExcelDateToTime(t)
		if convertTime.Second() != 0 {
			convertTime = convertTime.Add(time.Second * (time.Duration(60 - convertTime.Second())))
		}
		realTime = &convertTime
	}

	//}
	return realTime
}

// func ReadCurrentA() float32 {

// }

func ExcelDateToTime(serialDate float64) time.Time {
	// Excel uses 30-Dec-1899 as the epoch (0 date)
	excelEpoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)

	// Convert the serial number to a duration
	days := int(serialDate)
	seconds := (serialDate - float64(days)) * 86400 // seconds in a day

	// Add the days and seconds to the epoch
	finalTime := excelEpoch.AddDate(0, 0, days).Add(time.Duration(seconds) * time.Second)

	return finalTime
}
