package main

import (
	"log"
)

func main() {
	userCsvFile := "data/users.csv"
	numberOfWinners := 3

	records, err := ParseCsvFile(userCsvFile)
	if err != nil {
		log.Fatalf("Error parsing csv: %v", err)
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
		log.Printf("Winner: %v", winner)
	}
}
