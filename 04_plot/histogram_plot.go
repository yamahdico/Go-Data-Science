package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// داده‌های نمونه
	data := []float64{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 6}

	// ایجاد نمودار جدید
	p := plot.New()

	// عنوان و برچسب‌های محور
	p.Title.Text = "Histogram Example"
	p.X.Label.Text = "Value"
	p.Y.Label.Text = "Frequency"

	// ایجاد هیستوگرام و افزودن آن به نمودار
	hist, err := plotter.NewHist(plotter.Values(data), 16)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(hist)

	// ذخیره نمودار به فایل
	err = p.Save(4*vg.Inch, 4*vg.Inch, "histogram_plot.png")
	if err != nil {
		log.Fatal(err)
	}
}
