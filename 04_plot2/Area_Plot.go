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
		points[i].X = float64(i)
		points[i].Y = rand.Float64() * 100
	}

	// ایجاد نمودار
	p := plot.New()
	p.Title.Text = "Area Plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// افزودن ناحیه به نمودار
	area, err := plotter.NewArea(points)
	if err != nil {
		panic(err)
	}
	p.Add(area)

	// ذخیره نمودار
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "area_plot.png"); err != nil {
		panic(err)
	}
}
