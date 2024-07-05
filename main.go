package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	userCsvFile := flag.String("file", "data/large_users.csv", "Path to the user CSV file")
	numberOfWinners := flag.Int("winners", 6, "Number of winners to select")

	flag.Parse()

	records, err := ParseCsvFile(*userCsvFile)
	if err != nil {
		log.Fatalf("Error parsing CSV file: %v", err)
	}

	if len(records) <= 1 {
		log.Fatalf("Not enough records to select winners")
	}

	winners, err := SelectRandomEntries(records, *numberOfWinners)
	if err != nil {
		log.Fatalf("Error selecting winners: %v", err)
	}

	for _, winner := range winners {
		fmt.Printf("Winner: %s\n", winner)
	}
}
