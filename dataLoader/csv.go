package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
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
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// فرض می‌کنیم که فایل CSV دارای داده‌های عددی است
	// تبدیل داده‌ها به آرایه‌ای از اعداد (یک ماتریس 2D)
	var data [][]float64

	// پردازش رکوردهای CSV به نوع داده float64
	for _, record := range records {
		var row []float64
		for _, value := range record {
			// تبدیل هر مقدار به عدد
			num, err := strconv.ParseFloat(value, 64)
			if err == nil { // فقط داده‌های عددی را اضافه می‌کنیم
				row = append(row, num)
			}
		}
		data = append(data, row)
	}

	// تبدیل آرایه به یک ماتریس
	numRows := len(data)
	numCols := len(data[0])

	// انجام محاسبات آماری: محاسبه میانگین و انحراف معیار هر ستون
	fmt.Println("Column-wise Means:")
	for j := 0; j < numCols; j++ {
		var column []float64
		for i := 0; i < numRows; i++ {
			column = append(column, data[i][j])
		}
		mean := stat.Mean(column, nil)
		fmt.Printf("Column %d: Mean = %f\n", j+1, mean)
	}

	fmt.Println("\nColumn-wise Standard Deviations:")
	for j := 0; j < numCols; j++ {
		var column []float64
		for i := 0; i < numRows; i++ {
			column = append(column, data[i][j])
		}
		stdDev := stat.StdDev(column, nil)
		fmt.Printf("Column %d: Std Dev = %f\n", j+1, stdDev)
	}
}
