package main

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Create a new plot
	p := plot.New()

	// Set the title and axis labels
	p.Title.Text = "Sine Wave"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Create some data points for the sine wave
	n := 100
	pts := make(plotter.XYs, n)
	for i := range pts {
		x := float64(i) / 10.0
		pts[i].X = x
		pts[i].Y = math.Sin(x)
	}

	// Create a line plotter and add it to the plot
	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	// Save the plot to a PNG file
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "sine_wave.png"); err != nil {
		panic(err)
	}
}
