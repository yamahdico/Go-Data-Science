package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// داده‌های ستونی
	barWidth := 0.8
	values := plotter.Values{1, 2, 3, 4, 5}

	// ایجاد نمودار
	p := plot.New()
	p.Title.Text = "Bar Plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// افزودن نمودار ستونی به نمودار
	bars, err := plotter.NewBarChart(values, barWidth)
	if err != nil {
		panic(err)
	}
	p.Add(bars)

	// ذخیره نمودار
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "bar_plot.png"); err != nil {
		panic(err)
	}
}
