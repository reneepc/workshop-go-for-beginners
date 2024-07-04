package main

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"
)

func TestParseCsvFile(t *testing.T) {
	tests := map[string]struct {
		input       [][]string
		expected    [][]string
		setupError  func(t *testing.T, tmpDir string) string
		expectError bool
	}{
		"Empty CSV": {
			input:    [][]string{},
			expected: [][]string{},
		},
		"Single record": {
			input:    [][]string{{"Renê Cardozo", "rene.epcrdz@gmail.com"}},
			expected: [][]string{{"Renê Cardozo", "rene.epcrdz@gmail.com"}},
		},
		"Multiple records": {
			input: [][]string{
				{"Renê Cardozo", "rene.epcrdz@gmail.com"},
				{"Maria Silva", "maria.silva@example.com"},
				{"João Souza", "joao.souza@example.com"},
			},
			expected: [][]string{
				{"Renê Cardozo", "rene.epcrdz@gmail.com"},
				{"Maria Silva", "maria.silva@example.com"},
				{"João Souza", "joao.souza@example.com"},
			},
		},
		"Force os.Open error": {
			setupError: func(t *testing.T, tmpDir string) string {
				return filepath.Join(tmpDir, "non_existent.csv")
			},
			expectError: true,
		},
		"Force reader.ReadAll error": {
			setupError: func(t *testing.T, tmpDir string) string {
				malformedFile := filepath.Join(tmpDir, "malformed.csv")
				err := os.WriteFile(malformedFile, []byte(`"Name","Email"\n"Renê Cardozo"`), 0666)
				if err != nil {
					t.Fatalf("Failed to write malformed CSV file: %v", err)
				}
				return malformedFile
			},
			expectError: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tmpDir, err := os.MkdirTemp("", "test")
			if err != nil {
				t.Fatalf("Failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tmpDir)

			var inputFile string
			if tc.setupError != nil {
				inputFile = tc.setupError(t, tmpDir)
			} else {
				inputFile = filepath.Join(tmpDir, "input.csv")
				err = createTempCsvFile(t, inputFile, tc.input)
				if err != nil {
					t.Fatalf("Failed to create temp CSV file: %v", err)
				}
			}

			parsedRecords, err := ParseCsvFile(inputFile)
			if tc.expectError {
				if err == nil {
					t.Fatalf("Expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("ParseCsvFile() error: %v", err)
				}
				if !equal2DSlices(t, parsedRecords, tc.expected) {
					t.Errorf("ParseCsvFile() = %v, want %v", parsedRecords, tc.expected)
				}
			}
		})
	}
}

func TestSaveWinners(t *testing.T) {
	tests := map[string]struct {
		input       [][]string
		setupError  func(t *testing.T, tmpDir string) string
		expectError bool
	}{
		"Empty CSV": {
			input: [][]string{},
		},
		"Single record": {
			input: [][]string{{"Renê Cardozo", "rene.epcrdz@gmail.com"}},
		},
		"Multiple records": {
			input: [][]string{
				{"Renê Cardozo", "rene.epcrdz@gmail.com"},
				{"Maria Silva", "maria.silva@example.com"},
				{"João Souza", "joao.souza@example.com"},
			},
		},
		"Force os.Create error": {
			setupError: func(t *testing.T, tmpDir string) string {
				readOnlyDir := filepath.Join(tmpDir, "readonly")
				err := os.Mkdir(readOnlyDir, 0555)
				if err != nil {
					t.Fatalf("Failed to create read-only directory: %v", err)
				}
				return filepath.Join(readOnlyDir, "output.csv")
			},
			expectError: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tmpDir, err := os.MkdirTemp("", "test")
			if err != nil {
				t.Fatalf("Failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tmpDir)

			var outputFile string
			if tc.setupError != nil {
				outputFile = tc.setupError(t, tmpDir)
			} else {
				outputFile = filepath.Join(tmpDir, "output.csv")
			}

			err = SaveWinners(outputFile, tc.input)
			if tc.expectError {
				if err == nil {
					t.Fatalf("Expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("SaveWinners() error: %v", err)
				}
				savedRecords, err := ParseCsvFile(outputFile)
				if err != nil {
					t.Fatalf("Failed to read output CSV file: %v", err)
				}
				if !equal2DSlices(t, savedRecords, tc.input) {
					t.Errorf("SaveWinners() = %v, want %v", savedRecords, tc.input)
				}
			}
		})
	}
}

func createTempCsvFile(t *testing.T, filePath string, records [][]string) error {
	t.Helper()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}
