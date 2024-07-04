package main

import (
	"errors"
	"math/rand/v2"
)

func SelectRandomEntries(records [][]string, numberOfWinners int) ([][]string, error) {
	if numberOfWinners > len(records) {
		return nil, errors.New("Number of winners is greater than number of records")
	}

	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	return records[:numberOfWinners], nil
}

// OR

func SelectRandomEntries_Alt(records [][]string, numberOfWinners int) ([][]string, error) {
	if numberOfWinners > len(records) {
		return nil, errors.New("Number of winners is greater than number of records")
	}

	selected := make([][]string, numberOfWinners)
	for len(selected) < numberOfWinners {
		i := rand.IntN(len(records))
		selected = append(selected, records[i])
		records = append(records[:i], records[i+1:]...)
	}

	return selected, nil
}
