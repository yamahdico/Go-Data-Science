package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

func main() {
	// باز کردن فایل Excel
	xlFile, err := xlsx.OpenFile("example.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	// خواندن داده‌ها از شیت‌های مختلف
	for _, sheet := range xlFile.Sheets {
		fmt.Println("Reading sheet:", sheet.Name)
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Print(cell.String(), "\t")
			}
			fmt.Println()
		}
	}
}
