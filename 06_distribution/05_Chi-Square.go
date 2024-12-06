package main

import (
	"fmt"
	"math"
)

// Z-Score Calculation
func calculateZ(x, mu, sigma float64) float64 {
	return (x - mu) / sigma
}

// T-Score Calculation
func calculateT(xbar, mu, s float64, n int) float64 {
	return (xbar - mu) / (s / math.Sqrt(float64(n)))
}

// Chi-Square Calculation
func calculateChiSquare(observed, expected []float64) float64 {
	chiSquare := 0.0
	for i := range observed {
		chiSquare += math.Pow(observed[i]-expected[i], 2) / expected[i]
	}
	return chiSquare
}

func main() {
	// Example for Z-Score
	x, mu, sigma := 185.0, 170.0, 10.0
	zScore := calculateZ(x, mu, sigma)
	fmt.Printf("Z-Score: %.2f\n", zScore)

	// Example for T-Score
	xbar, mu, s := 85.0, 90.0, 10.0
	n := 15
	tScore := calculateT(xbar, mu, s, n)
	fmt.Printf("T-Score: %.2f\n", tScore)

	// Example for Chi-Square
	observed := []float64{12, 18, 28}
	expected := []float64{10, 20, 30}
	chiSquare := calculateChiSquare(observed, expected)
	fmt.Printf("Chi-Square: %.3f\n", chiSquare)
}
