package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// داده‌های تصادفی برای جعبه‌ها
	data := [][]float64{
		randomData(),
		randomData(),
		randomData(),
	}

	// ایجاد نمودار
	p := plot.New()
	p.Title.Text = "Box Plot"
	p.Y.Label.Text = "Value"

	// آماده‌سازی داده‌ها برای جعبه‌ها
	boxes := make([]plotter.Values, len(data))
	for i, d := range data {
		boxes[i] = plotter.Values(d)
	}

	// افزودن جعبه‌ها به نمودار
	for _, b := range boxes {
		p.Add(plotter.NewBoxPlot(b))
	}

	// ذخیره نمودار
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "box_plot.png"); err != nil {
		panic(err)
	}
}

func randomData() []float64 {
	data := make([]float64, 20)
	for i := range data {
		data[i] = rand.Float64() * 100
	}
	return data
}
