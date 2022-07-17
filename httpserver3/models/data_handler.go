package models

import (
	"encoding/csv"
	"log"
	"os"
)

var file_path string

func ReadData(fileName string) ([][]string, error) {

	f, err := os.Open(file_path + fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	// if _, err := r.Read(); err != nil {
	// 	return [][]string{}, err
	// }

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func WriteData(fileName string, record []string) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if os.IsNotExist(err) {
		log.Println("file not exist, now try to crete")
		error := os.MkdirAll(file_path, os.ModePerm)
		if error != nil {
			log.Println(error)
		}
		f, err = os.Create(fileName)

	}
	if err != nil {
		log.Println("failed to open file", err)
	}

	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if err := w.Write(record); err != nil {
		log.Println("error writing record to file", err)
	}
}
