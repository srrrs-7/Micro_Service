package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func createCsv(headers []string, values [][]string) {
	// file name
	f, err := createCsvFile()
	if err != nil {
		log.Fatalf("cannot create csv file  : %s", err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	writer.Write(headers)

	err = writer.WriteAll(values)
	if err != nil {
		log.Fatalf("csv writer error : %s", err)
	}
}

func createCsvFile() (*os.File, error) {
	filePath := "./output/"
	now := time.Now()
	fileName := fmt.Sprintf("%d-%d-%d%d:%d:%d.csv",
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
	)
	// create csv file
	f, err := os.Create(filePath + fileName)
	if err != nil {
		return nil, err
	}

	return f, nil
}
