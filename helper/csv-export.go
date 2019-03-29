package helper

import (
	"encoding/csv"
	"log"
	"os"
	"time"
	"bitbucket.org/babulal107/go-app/config"
)

func GenerateCSV(fileName string, data [][]string) error{
	var (
		path string
		writer *csv.Writer
		file *os.File
		err error
	)

	path = config.FileExportPath+GetFileName(fileName)+config.FileExtenuationCST
	if file, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm); err!=nil{
		log.Fatal("Cannot create file", err)
		return err
	}
	defer Close(file)

	defer func() {
		if writer != nil {
			writer.Flush()
		}
	}()

	writer = csv.NewWriter(file)
	if err = writer.WriteAll(data); err!=nil{
		log.Fatal("Cannot write to file", err)
		return err
	}
	return err
}

func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func GetFileName(fileName string) string{
	var timeStamp = time.Now().Format(config.TimeStampFormat)
	return fileName+"_"+timeStamp
}