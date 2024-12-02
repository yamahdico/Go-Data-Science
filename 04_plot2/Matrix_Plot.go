package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// داده‌های ماتریسی (مثال)
	matrix := make([][]float64, 10)
	for i := range matrix {
		matrix[i] = make([]float64, 10)
		for j := range matrix[i] {
			matrix[i][j] = float64(i * j)
		}
	}

	// ایجاد نمودار
	p := plot.New()
	p.Title.Text = "Matrix Plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// افزودن ماتریس به نمودار
	heatMap := plotter.NewHeatMap(matrix, palette.Heat(256))
	p.Add(heatMap)

	// ذخیره نمودار
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "matrix_plot.png"); err != nil {
		panic(err)
	}
}
