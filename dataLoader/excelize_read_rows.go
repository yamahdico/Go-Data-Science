package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	// باز کردن فایل Excel
	f, err := excelize.OpenFile("dataLoader/example.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	// خواندن داده‌ها از شیت مورد نظر (مثلاً "Sheet1")
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	// نمایش داده‌ها
	for _, row := range rows {
		fmt.Println(row)
	}
}
