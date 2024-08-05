package main

import (
	"fmt"
	"testing"

	"math/rand/v2"
)

var recordSizes = []int{100, 1000, 10000}
var numberOfWinners = 10
var recordLength = 10

func BenchmarkSelectRandomEntries(b *testing.B) {
	for _, recordSize := range recordSizes {
		records := generateBenchmarkRecords(recordSize, recordLength)
		testName := fmt.Sprintf("SelectRandomEntries/%d", recordSize)
		b.Run(testName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := SelectRandomEntries(records, numberOfWinners)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkSelectRandomEntries_Alt(b *testing.B) {
	for _, recordSize := range recordSizes {
		records := generateBenchmarkRecords(recordSize, recordLength)
		testName := fmt.Sprintf("SelectRandomEntries_Alt/%d", recordSize)
		b.Run(testName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := SelectRandomEntries_Alt(records, numberOfWinners)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkSelectRandomEntries_WithoutCopying(b *testing.B) {
	for _, recordSize := range recordSizes {
		records := generateBenchmarkRecords(recordSize, recordLength)
		testName := fmt.Sprintf("SelectRandomEntries_WithoutCopying/%d", recordSize)
		b.Run(testName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := SelectRandomEntries_WithoutCopying(records, numberOfWinners)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func generateBenchmarkRecords(numRecords, recordLength int) [][]string {
	records := make([][]string, numRecords)
	for i := 0; i < numRecords; i++ {
		record := make([]string, recordLength)
		for j := 0; j < recordLength; j++ {
			record[j] = randomStringOfNumbers(5)
		}
		records[i] = record
	}
	return records
}

func randomStringOfNumbers(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = byte('0' + rand.IntN(10))
	}
	return string(b)
}
