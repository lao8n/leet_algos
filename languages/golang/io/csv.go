package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("sample.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a new reader.
	reader := csv.NewReader(file)

	// Read all the records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the records
	for i, record := range records {
		// Skip header row
		if i == 0 {
			continue
		}

		fmt.Printf("Name: %s, Age: %s, Email: %s\n", record[0], record[1], record[2])
	}
}
