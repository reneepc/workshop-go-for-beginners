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

	winners, err := SelectRandomEntries(records, numberOfWinners)
	if err != nil {
		log.Fatalf("Error selecting winners: %v", err)
	}

}
