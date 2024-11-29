package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	// باز کردن فایل اکسل
	f, err := excelize.OpenFile("dataLoader/example.xlsx")
	if err != nil {
		log.Fatalf("Failed to open Excel file: %v", err)
	}
	defer f.Close()

	// خواندن سطر خاص (مثلاً سطر شماره 2)
	rowNumber := 2
	sheetName := "Sheet1" // نام شیت مورد نظر
	row, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Failed to get rows: %v", err)
	}

	if rowNumber <= len(row) {
		fmt.Printf("Row %d: %v\n", rowNumber, row[rowNumber-1])
	} else {
		fmt.Printf("Row %d does not exist\n", rowNumber)
	}
}
