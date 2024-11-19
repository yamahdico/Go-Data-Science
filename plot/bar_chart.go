package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// داده‌های نمونه
	values := []float64{1, 2, 3, 4, 5}

	// ایجاد نمودار جدید
	p := plot.New()

	// عنوان و برچسب‌های محور
	p.Title.Text = "Bar Chart Example"
	p.X.Label.Text = "Category"
	p.Y.Label.Text = "Value"

	// ایجاد داده‌های میله‌ای
	bars := make(plotter.Values, len(values))
	for i, v := range values {
		bars[i] = v
	}

	// ایجاد نمودار میله‌ای
	bar, err := plotter.NewBarChart(bars, vg.Points(20))
	if err != nil {
		log.Fatal(err)
	}
	p.Add(bar)

	// ذخیره نمودار به فایل
	err = p.Save(4*vg.Inch, 4*vg.Inch, "bar_chart.png")
	if err != nil {
		log.Fatal(err)
	}
}
