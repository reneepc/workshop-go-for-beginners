package main

import (
	"fmt"
	"log"
)

func main() {
	userCsvFile := "data/large_users_unfair.csv"
	numberOfWinners := 6

	records, err := ParseCsvFile(userCsvFile)
	if err != nil {
		log.Fatalf("Error parsing CSV file: %v", err)
	}

	if len(records) <= 1 {
		log.Fatalf("Not enough records to select winners")
	}

	if numberOfWinners > len(records) {
		log.Fatalf("Number of winners is greater than number of records")
	}

	winners, err := SelectRandomEntries(records, numberOfWinners)
	if err != nil {
		log.Fatalf("Error selecting winners: %v", err)
	}

	for _, winner := range winners {
		fmt.Printf("Winner: %s\n", winner)
	}
}
