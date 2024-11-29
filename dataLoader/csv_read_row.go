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

	for i := 0; i < 5; i++ {
		row, err := reader.Read()
		if err != nil {
			break // اگر به انتهای فایل رسیدیم، حلقه متوقف شود
		}
		fmt.Println(row)
	}

	if err != nil {
		log.Fatal(err)
	}

}
