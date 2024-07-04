package main

import (
	"slices"
	"testing"
)

func TestSelectRandomEntries(t *testing.T) {
	records := [][]string{
		{"Renê Cardozo", "rene.epcrdz@gmail.com"},
		{"Maria Silva", "maria.silva@example.com"},
		{"João Souza", "joao.souza@example.com"},
		{"Ana Pereira", "ana.pereira@example.com"},
		{"Carlos Lima", "carlos.lima@example.com"},
	}

	t.Run("Basic test", func(t *testing.T) {
		numberOfWinners := 3

		winners, err := SelectRandomEntries(records, numberOfWinners)
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		if len(winners) != numberOfWinners {
			t.Fatalf("Expected %d winners, but got %d", numberOfWinners, len(winners))
		}

		for _, winner := range winners {
			if !containsRecord(t, records, winner) {
				t.Errorf("Winners not found in original records")
			}
		}
	})

	t.Run("Empty records", func(t *testing.T) {
		emptyRecords := [][]string{}
		_, err := SelectRandomEntries(emptyRecords, 3)
		if err == nil {
			t.Fatalf("Expected error for empty records, but got nil")
		}
	})

	t.Run("Zero winners", func(t *testing.T) {
		numberOfWinners := 0

		winners, err := SelectRandomEntries(records, numberOfWinners)
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		if len(winners) != numberOfWinners {
			t.Fatalf("Expected %d winners, but got %d", numberOfWinners, len(winners))
		}
	})

	t.Run("Negative number of winners", func(t *testing.T) {
		_, err := SelectRandomEntries(records, -1)
		if err == nil {
			t.Fatalf("Expected error for negative number of winners, but got nil")
		}
	})

	t.Run("Number of winners greater than number of records", func(t *testing.T) {
		_, err := SelectRandomEntries(records, len(records)+1)
		if err == nil {
			t.Fatalf("Expected error for number of winners greater than number of records, but got nil")
		}
	})
}

func TestSelectRandomEntriesAlt(t *testing.T) {
	records := [][]string{
		{"Renê Cardozo", "rene.epcrdz@gmail.com"},
		{"Maria Silva", "maria.silva@example.com"},
		{"João Souza", "joao.souza@example.com"},
		{"Ana Pereira", "ana.pereira@example.com"},
		{"Carlos Lima", "carlos.lima@example.com"},
	}

	t.Run("Basic test", func(t *testing.T) {
		numberOfWinners := 3

		winners, err := SelectRandomEntries_Alt(records, numberOfWinners)
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		if len(winners) != numberOfWinners {
			t.Fatalf("Expected %d winners, but got %d", numberOfWinners, len(winners))
		}

		for _, winner := range winners {
			if !containsRecord(t, records, winner) {
				t.Errorf("Winners not found in original records")
			}
		}
	})

	t.Run("Empty records", func(t *testing.T) {
		emptyRecords := [][]string{}
		_, err := SelectRandomEntries_Alt(emptyRecords, 3)
		if err == nil {
			t.Fatalf("Expected error for empty records, but got nil")
		}
	})

	t.Run("Zero winners", func(t *testing.T) {
		numberOfWinners := 0

		winners, err := SelectRandomEntries_Alt(records, numberOfWinners)
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		if len(winners) != numberOfWinners {
			t.Fatalf("Expected %d winners, but got %d", numberOfWinners, len(winners))
		}
	})

	t.Run("Negative number of winners", func(t *testing.T) {
		_, err := SelectRandomEntries_Alt(records, -1)
		if err == nil {
			t.Fatalf("Expected error for negative number of winners, but got nil")
		}
	})

	t.Run("Number of winners greater than number of records", func(t *testing.T) {
		_, err := SelectRandomEntries_Alt(records, len(records)+1)
		if err == nil {
			t.Fatalf("Expected error for number of winners greater than number of records, but got nil")
		}
	})
}

func containsRecord(t *testing.T, records [][]string, record []string) bool {
	t.Helper()

	for _, r := range records {
		if slices.Equal(r, record) {
			return true
		}
	}
	return false
}
