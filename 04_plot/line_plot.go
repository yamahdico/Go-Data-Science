package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// ایجاد نمودار جدید
	p := plot.New()
	p.Title.Text = "Line Plot Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// داده‌های نمودار
	pts := plotter.XYs{
		{X: 1, Y: 2},
		{X: 2, Y: 4},
		{X: 3, Y: 8},
	}

	// اضافه کردن داده‌ها به نمودار
	line, _ := plotter.NewLine(pts)
	p.Add(line)

	// ذخیره به فایل
	p.Save(4*vg.Inch, 4*vg.Inch, "line_plot.png")
}
