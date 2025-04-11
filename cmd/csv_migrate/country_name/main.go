package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open CSV file
	file, err := os.Open("./port_location.csv")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}

	// Connect to MySQL
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/port")
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	defer db.Close()

	// Assuming first row is header
	for i, row := range records {
		if i == 0 {
			continue
		}
		portCode := row[0]
		countryCode := row[2]
		countryNameCN := row[3]
		countryNameEN := row[4]

		_, err := db.Exec(`
            UPDATE port 
            SET countryNameCN = ?, countryNameEN = ?, countryCode = ?
            WHERE portCode = ?
        `, countryNameCN, countryNameEN, countryCode, portCode)

		if err != nil {
			log.Printf("Failed to update port %s: %v", portCode, err)
		} else {
			fmt.Printf("Updated portCode %s with CN: %s, EN: %s\n", portCode, countryNameCN, countryNameEN)
		}
	}

	fmt.Println("Update completed.")
}
