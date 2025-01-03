package filedata

import "htmxll/repository"

type FileData interface {
	CheckNewFileRealTime()
	InitReadFile()
}

type fileData struct {
	dataTempRepo repository.Repository
}

func NewFileData(dataTempRepo repository.Repository) FileData {
	return fileData{dataTempRepo}
}
