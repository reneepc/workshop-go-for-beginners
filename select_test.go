package main

import (
	"encoding/json"
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

func FuzzSelectRandomEntriesAlt(f *testing.F) {
	initialRecords := [][]string{
		{"Renê Cardozo", "rene.epcrdz@gmail.com"},
		{"Maria Silva", "maria.silva@example.com"},
		{"João Souza", "joao.souza@example.com"},
		{"Ana Pereira", "ana.pereira@example.com"},
		{"Carlos Lima", "carlos.lima@example.com"},
	}
	initialData, _ := json.Marshal(initialRecords)
	f.Add(string(initialData), 3)

	f.Fuzz(func(t *testing.T, recordsJSON string, numberOfWinners int) {
		var records [][]string
		if err := json.Unmarshal([]byte(recordsJSON), &records); err != nil {
			return
		}
		if len(records) == 0 {
			return
		}

		winners, err := SelectRandomEntries(records, numberOfWinners)

		if numberOfWinners < 0 || numberOfWinners > len(records) {
			if err == nil {
				t.Errorf("Expected error for invalid number of winners, but got nil")
			}
		} else {
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
		}
	})
}

func TestRemoveDuplicates(t *testing.T) {
	tests := map[string]struct {
		input    [][]string
		expected [][]string
	}{
		"No duplicates": {
			input:    [][]string{{"Renê Cardozo", "rene.epcrdz@gmail.com"}, {"Maria Silva", "maria.silva@example.com"}},
			expected: [][]string{{"Renê Cardozo", "rene.epcrdz@gmail.com"}, {"Maria Silva", "maria.silva@example.com"}},
		},
		"All duplicates": {
			input: [][]string{{"Renê Cardozo", "rene.epcrdz@gmail.com"}, {"Renê Cardozo", "rene.epcrdz@gmail.com"},
				{"Renê Cardozo", "rene.epcrdz@gmail.com"}, {"Renê Cardozo", "rene.epcrdz@gmail.com"},
				{"Renê Cardozo", "rene.epcrdz@gmail.com"}},
			expected: [][]string{{"Renê Cardozo", "rene.epcrdz@gmail.com"}},
		},
		"One duplicates": {
			input:    [][]string{{"Renê Cardozo", "rene.epcrdz@gmail.com"}, {"Maria Silva", "maria.silva@example.com"}, {"Renê Cardozo", "rene.epcrdz@gmail.com"}},
			expected: [][]string{{"Renê Cardozo", "rene.epcrdz@gmail.com"}, {"Maria Silva", "maria.silva@example.com"}},
		},
		"Empty input": {
			input:    [][]string{},
			expected: [][]string{},
		},
		"Single record": {
			input:    [][]string{{"Renê Cardozo", "rene.epcrdz@gmail.com"}},
			expected: [][]string{{"Renê Cardozo", "rene.epcrdz@gmail.com"}},
		},
		"Multiple duplicates": {
			input: [][]string{
				{"Renê Cardozo", "rene.epcrdz@gmail.com"},
				{"Maria Silva", "maria.silva@example.com"},
				{"João Souza", "joao.souza@example.com"},
				{"Renê Cardozo", "rene.epcrdz@gmail.com"},
				{"Maria Silva", "maria.silva@example.com"},
				{"Ana Pereira", "ana.pereira@example.com"},
				{"João Souza", "joao.souza@example.com"},
			},
			expected: [][]string{
				{"Renê Cardozo", "rene.epcrdz@gmail.com"},
				{"Maria Silva", "maria.silva@example.com"},
				{"João Souza", "joao.souza@example.com"},
				{"Ana Pereira", "ana.pereira@example.com"},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := RemoveDuplicates(tt.input)
			if !equal2DSlices(t, result, tt.expected) {
				t.Errorf("RemoveDuplicates() = %v, want %v", result, tt.expected)
			}
		})
	}
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

func equal2DSlices(t *testing.T, a, b [][]string) bool {
	t.Helper()

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !slices.Equal(a[i], b[i]) {
			return false
		}
	}

	return true
}
