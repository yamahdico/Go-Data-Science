package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// داده‌های تصادفی
	n := 1000
	values := make(plotter.Values, n)
	for i := 0; i < n; i++ {
		values[i] = rand.NormFloat64()
	}

	// ایجاد نمودار
	p := plot.New()
	p.Title.Text = "Density Plot (Histogram)"
	p.X.Label.Text = "Value"
	p.Y.Label.Text = "Frequency"

	// افزودن هیستوگرام به نمودار
	h, err := plotter.NewHist(values, 16)
	if err != nil {
		panic(err)
	}
	p.Add(h)

	// ذخیره نمودار
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "density_plot.png"); err != nil {
		panic(err)
	}
}
