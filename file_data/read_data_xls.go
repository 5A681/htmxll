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
				dataTempRepo.CreateBay(&entity.Bay{
					Name:         ws.Name,
					SubStationId: sub.Id,
					SheetNumber:  sheet,
				})
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
	if ws.Row(2).Col(2) == "VOLTAGE PHASE B-C" {
		for r := 5; r < maxRow; r++ {

			if ws.Row(r).Col(0) != "" {

				tempData := entity.DataTmps{
					BayId: bay.Id,
				}
				for c := 0; c < maxCol; c++ {

					if c == 0 {

						//fmt.Printf("%v \t", ReadDateTimeColumn(ws.Row(r).Col(c)))
						dateTime := ReadDateTimeColumn(ws.Row(r).Col(c))
						if dateTime != nil {
							tempData.DataDatetime = *dateTime
						}
					}

					if c == 2 || c == 4 || c == 6 || c == 8 || c == 10 || c == 12 || c == 14 {
						//fmt.Printf("%v \t", ws.Row(r).Col(c))
						if c > 14 {
							continue
						} else if c == 2 {

							floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
							if err != nil {
								tempData.VoltageBC = 0
							} else {
								tempData.VoltageBC = float32(floatData)
							}
							if ws.Name == "INCOMING1" {
							}

						} else if c == 4 {
							floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
							if err != nil {
								tempData.CurrentPhaseA = 0
							} else {
								tempData.CurrentPhaseA = float32(floatData)
							}
						} else if c == 6 {
							floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
							if err != nil {
								tempData.CurrentPhaseB = 0
							} else {
								tempData.CurrentPhaseB = float32(floatData)
							}
						} else if c == 8 {
							floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
							if err != nil {
								tempData.CurrentPhaseC = 0

							} else {
								tempData.CurrentPhaseC = float32(floatData)

							}
						} else if c == 10 {
							floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
							if err != nil {
								tempData.ActivePower = 0

							} else {
								tempData.ActivePower = float32(floatData)
							}
						} else if c == 12 {
							floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
							if err != nil {
								tempData.ReactivePower = 0
							} else {
								tempData.ReactivePower = float32(floatData)
							}
						} else if c == 14 {
							floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
							if err != nil {
								tempData.PowerFactor = 0

							} else {
								tempData.PowerFactor = float32(floatData)
							}
						}
					} else {
						continue
					}

				}
				tempData.CreatedAt = time.Now()
				err := dataTempRepo.CreateDataTmep(tempData)
				if err != nil {
					log.Println("could not insrt temp data", err.Error())
				}
			}

			//fmt.Printf("\n\n")
		}
	} else {
		for r := 5; r < maxRow; r++ {

			if ws.Row(r).Col(0) != "" {

				tempData := entity.DataTmps{
					BayId: bay.Id,
				}
				for c := 0; c < maxCol; c++ {

					if c == 0 {

						//fmt.Printf("%v \t", ReadDateTimeColumn(ws.Row(r).Col(c)))
						dateTime := ReadDateTimeColumn(ws.Row(r).Col(c))
						if dateTime != nil {
							tempData.DataDatetime = *dateTime
						}
					}

					if c == 2 || c == 4 || c == 6 || c == 8 || c == 10 || c == 12 {
						//fmt.Printf("%v \t", ws.Row(r).Col(c))
						if c > 12 {
							continue
						} else if c == 2 {
							floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
							if err != nil {
								tempData.CurrentPhaseA = 0
							} else {
								tempData.CurrentPhaseA = float32(floatData)
							}
						} else if c == 4 {
							floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
							if err != nil {
								tempData.CurrentPhaseB = 0
							} else {
								tempData.CurrentPhaseB = float32(floatData)
							}
						} else if c == 6 {
							floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
							if err != nil {
								tempData.CurrentPhaseC = 0
							} else {
								tempData.CurrentPhaseC = float32(floatData)
							}
						} else if c == 8 {
							floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
							if err != nil {
								tempData.ActivePower = 0
							} else {
								tempData.ActivePower = float32(floatData)
							}
						} else if c == 10 {
							floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
							if err != nil {
								tempData.ReactivePower = 0
							} else {
								tempData.ReactivePower = float32(floatData)
							}
						} else if c == 12 {
							floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
							if err != nil {
								tempData.PowerFactor = 0
							} else {
								tempData.PowerFactor = float32(floatData)
							}
						}
					} else {
						continue
					}

				}
				tempData.CreatedAt = time.Now()

				err := dataTempRepo.CreateDataTmep(tempData)
				if err != nil {
					log.Println("could not insrt temp data", err.Error())
				}
			}

			//fmt.Printf("\n\n")
		}
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
