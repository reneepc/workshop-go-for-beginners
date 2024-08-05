package main

import (
	"errors"
	"math/rand/v2"
)

func SelectRandomEntries(records [][]string, numberOfWinners int) ([][]string, error) {
	if numberOfWinners > len(records) {
		return nil, errors.New("number of winners greater than number of records")
	}

	if numberOfWinners < 0 {
		return nil, errors.New("negative number of winners")
	}

	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	return records[:numberOfWinners], nil
}

// OR

func SelectRandomEntries_Alt(records [][]string, numberOfWinners int) ([][]string, error) {
	if numberOfWinners > len(records) {
		return nil, errors.New("number of winners greater than number of records")
	}

	if numberOfWinners < 0 {
		return nil, errors.New("negative number of winners")
	}

	recordsCopy := make([][]string, len(records))
	copy(recordsCopy, records)

	selected := make([][]string, 0, numberOfWinners)
	for len(selected) < numberOfWinners {
		i := rand.IntN(len(recordsCopy))
		selected = append(selected, recordsCopy[i])
		recordsCopy = append(recordsCopy[:i], recordsCopy[i+1:]...)
	}

	return selected, nil
}

func SelectRandomEntries_WithoutCopying(records [][]string, numberOfWinners int) ([][]string, error) {
	if numberOfWinners > len(records) {
		return nil, errors.New("number of winners greater than number of records")
	}

	if numberOfWinners < 0 {
		return nil, errors.New("negative number of winners")
	}

	selected := make([][]string, 0, numberOfWinners)
	for len(selected) < numberOfWinners {
		i := rand.IntN(len(records))
		selected = append(selected, records[i])
		records = append(records[:i], records[i+1:]...)
	}

	return selected, nil
}

func RemoveDuplicates(records [][]string) [][]string {
	uniqueRecords := make(map[string]bool)
	deduplicated := [][]string{}

	for _, record := range records {
		email := record[0]
		if _, exists := uniqueRecords[email]; !exists {
			uniqueRecords[email] = true
			deduplicated = append(deduplicated, record)
		}
	}

	return deduplicated
}
