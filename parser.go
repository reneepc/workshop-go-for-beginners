package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ParseCsvFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV file: %v", err)
	}

	return records, nil
}

func SaveWinners(filePath string, winners [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, winner := range winners {
		if err := writer.Write(winner); err != nil {
			return fmt.Errorf("error writing winner to CSV file: %v", err)
		}
	}
	return nil
}
