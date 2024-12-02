package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"math/rand"
)

func main() {
	// داده‌ها برای دو محور
	n := 100
	x := make([]float64, n)
	y1 := make([]float64, n)
	y2 := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = float64(i)
		y1[i] = rand.Float64() * 100
		y2[i] = rand.Float64() * 50
	}

	// ایجاد نمودار
	p := plot.New()
	p.Title.Text = "Dual Axis Plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y1"

	// نمودار اول (Y1)
	line1, err := plotter.NewLine(plotter.XYs{X: x, Y: y1})
	if err != nil {
		panic(err)
	}
	p.Add(line1)

	// ایجاد محور دوم
	p2 := plot.New()
	p2.Y.Label.Text = "Y2"
	line2, err := plotter.NewLine(plotter.XYs{X: x, Y: y2})
	if err != nil {
		panic(err)
	}
	p2.Add(line2)

	// ذخیره نمودار
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "dual_axis_plot.png"); err != nil {
		panic(err)
	}
}
