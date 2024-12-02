package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// داده‌های حبابی
	data := []plotter.XY{
		{X: 10, Y: 20},
		{X: 15, Y: 25},
		{X: 20, Y: 30},
	}

	// ایجاد نمودار
	p := plot.New()
	p.Title.Text = "Bubble Chart"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// ایجاد پراکندگی با اندازه حباب‌ها
	bubbles := plotter.NewScatter(data)
	bubbles.GlyphStyle.Radius = vg.Points(10) // اندازه حباب‌ها
	p.Add(bubbles)

	// ذخیره نمودار
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "bubble_chart.png"); err != nil {
		panic(err)
	}
}
