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

	// تحلیل داده‌ها (مثال ساده)
	fmt.Println("Response data:", data)
}
