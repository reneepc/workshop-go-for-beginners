package main

import (
	"path/filepath"
	"slices"
	"testing"
)

func TestSelectRandomUsers(t *testing.T) {
	testDataDir := "testdata"
	filePattern := "*.csv"
	numberOfWinners := 8

	files, err := filepath.Glob(filepath.Join(testDataDir, filePattern))
	if err != nil {
		t.Fatalf("Failed to find test files: %v", err)
	}

	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			records, err := ParseCsvFile(file)
			if err != nil {
				t.Fatalf("Failed to parse CSV file %s: %v", file, err)
			}

			records = RemoveDuplicates(records)

			if len(records) < numberOfWinners {
				t.Fatalf("File %s has fewer records than number of winners", file)
			}

			winners, err := SelectRandomEntries(records, numberOfWinners)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}

			if len(winners) != numberOfWinners {
				t.Errorf("Expected %d winners, got %d", numberOfWinners, len(winners))
			}

			for _, winner := range winners {
				if !contains(records, winner) {
					t.Errorf("Winner %v not found in original records", winner)
				}
			}

			if hasDuplicatesBasedOnFirstEntry(winners) {
				t.Log(winners)
				t.Errorf("Duplicate winners found: %v", winners)
			}
		})
	}
}

func contains(records [][]string, winner []string) bool {
	return slices.ContainsFunc(records, func(record []string) bool {
		return slices.Equal(record, winner)
	})
}

func hasDuplicatesBasedOnFirstEntry(records [][]string) bool {
	seen := make(map[string]struct{})
	for _, record := range records {
		if len(record) == 0 {
			continue // Skip empty records
		}
		email := record[0]
		if _, exists := seen[email]; exists {
			return true
		}
		seen[email] = struct{}{}
	}
	return false
}
