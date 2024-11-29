package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// باز کردن فایل CSV
	file, err := os.Open("dataLoader/example.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// خواندن داده‌ها از CSV
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// نمایش ستون اول و سوم
	fmt.Println("Column 1 and Column 3:")
	for _, row := range data {
		if len(row) > 2 { // بررسی وجود حداقل سه ستون
			fmt.Printf("Col1: %s, Col3: %s\n", row[0], row[2])
		}
	}
}
