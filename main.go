package main

import (
	"flag"
	"log"
	"os"

	"github.com/go-gomail/gomail"
)

func main() {
	userCsvFile := flag.String("file", "testdata/users_001.csv", "Path to the user CSV file")
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

	if err := SendEmail(winners); err != nil {
		log.Fatalf("Error sending email to winners")
	}

	if err := SaveWinners("winners.csv", winners); err != nil {
		log.Fatalf("Error saving winners to csv")
	}
}

func SendEmail(winners [][]string) error {
	email := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	dialer := gomail.NewDialer("smtp.gmail.com", 587, email, password)

	emails := extractEmails(winners)

	message := gomail.NewMessage()
	message.SetHeader("Subject", "You are a winner!")
	message.SetHeader("From", email)
	message.SetHeader("To", emails...)

	htmlString, err := os.ReadFile("email.html")
	if err != nil {
		log.Fatalf("Error reading email template: %v", err)
	}

	message.SetBody("text/html", string(htmlString))

	if err := dialer.DialAndSend(message); err != nil {
		log.Fatalf("Error sending email: %v", err)
	}

	return nil
}

func extractEmails(winners [][]string) []string {
	emails := make([]string, 0, len(winners))
	for _, winner := range winners {
		emails = append(emails, winner[0])
	}

	return emails
}
