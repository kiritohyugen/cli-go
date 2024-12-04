package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generateCSV(fileName string, rows int) error {
	// Create or open the CSV file
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	err = writer.Write([]string{"Col1", "Col2"})
	if err != nil {
		return fmt.Errorf("could not write header: %w", err)
	}

	// Write the data rows
	for i := 0; i < rows; i++ {
		// Generate random data for "Col1" and "Col2"
		col1 := "Data" + strconv.Itoa(i)
		col2 := rand.Intn(20001) - 10000 // Random number between -10000 and 10000

		// Write the row
		err := writer.Write([]string{col1, strconv.Itoa(col2)})
		if err != nil {
			return fmt.Errorf("could not write row: %w", err)
		}
	}

	return nil
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Number of files (m) and rows (n)
	m := 5  // Example: creating 5 CSV files
	n := 10 // Example: each file will have 10 rows

	// Create m CSV files, each with n rows
	for i := 1; i <= m; i++ {
		fileName := fmt.Sprintf("file%d.csv", i)
		err := generateCSV(fileName, n)
		if err != nil {
			fmt.Printf("Error creating %s: %v\n", fileName, err)
		} else {
			fmt.Printf("%s created successfully!\n", fileName)
		}
	}
}
