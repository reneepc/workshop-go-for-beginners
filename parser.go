package main

import (
	"encoding/csv"
	"log"
	"os"
)

func ParseCsvFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading csv: %v", err)
	}

	return records, nil
}

func SaveWinners(filePath string, winners [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, winner := range winners {
		if err := writer.Write(winner); err != nil {
			return err
		}
	}
	return nil
}
