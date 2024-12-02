package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// ایجاد نمودار جدید
	p := plot.New()

	// عنوان و برچسب‌های محور
	p.Title.Text = "Scatter Plot Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// داده‌ها
	pts := plotter.XYs{
		{X: 1, Y: 2},
		{X: 2, Y: 3},
		{X: 3, Y: 5},
		{X: 4, Y: 7},
		{X: 5, Y: 11},
	}

	// ایجاد نمودار پراکندگی و افزودن آن به نمودار
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(scatter)

	// ذخیره نمودار به فایل
	err = p.Save(4*vg.Inch, 4*vg.Inch, "scatter_plot.png")
	if err != nil {
		log.Fatal(err)
	}
}
