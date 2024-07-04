package main

import (
	"math/rand/v2"
)

func SelectRandomEntries(records [][]string, numberOfWinners int) ([][]string, error) {
	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	return records[:numberOfWinners], nil
}

// OR

func SelectRandomEntries_Alt(records [][]string, numberOfWinners int) ([][]string, error) {
	selected := make([][]string, numberOfWinners)
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
