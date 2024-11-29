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

	// خواندن سرستون‌ها از CSV
	reader := csv.NewReader(file)
	headers, err := reader.Read() // فقط اولین سطر (Header)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Headers:", headers)
}
