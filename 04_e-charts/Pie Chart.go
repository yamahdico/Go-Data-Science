package main

import (
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	// داده‌های نمونه برای نمودار دایره‌ای
	data := []opts.PieData{
		{Name: "Category A", Value: 40},
		{Name: "Category B", Value: 30},
		{Name: "Category C", Value: 20},
		{Name: "Category D", Value: 10},
	}

	// ایجاد نمودار دایره‌ای
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Pie Chart Example"}),
	)
	pie.AddSeries("Categories", data)

	// ذخیره نمودار به فایل HTML
	f, err := os.Create("pie_chart.html")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer f.Close()

	if err := pie.Render(f); err != nil {
		log.Fatalf("Error rendering chart: %v", err)
	}
}
