package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// داده‌های تصادفی
	n := 100
	points := make(plotter.XYs, n)
	for i := 0; i < n; i++ {
		points[i].X = rand.Float64() * 10
		points[i].Y = rand.Float64() * 10
	}

	// ایجاد نمودار
	p := plot.New()
	p.Title.Text = "Scatter Plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// افزودن نقاط به نمودار
	scatter, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	p.Add(scatter)

	// ذخیره نمودار
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "scatter_plot.png"); err != nil {
		panic(err)
	}
}
