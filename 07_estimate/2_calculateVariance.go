package main

import (
	"fmt"
	"math"
)

// calculateMean calculates the sample mean
func calculateMean(data []float64) float64 {
	var sum float64
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// calculateVariance calculates the sample variance
func calculateVariance(data []float64, mean float64) float64 {
	var sumOfSquares float64
	for _, value := range data {
		sumOfSquares += math.Pow(value-mean, 2)
	}
	return sumOfSquares / float64(len(data)-1)
}

func main() {
	// Sample data
	data := []float64{70, 75, 80, 85, 90}

	// Calculate the sample mean
	mean := calculateMean(data)
	fmt.Printf("Point estimate of the population mean: %.2f\n", mean)

	// Calculate the sample variance
	variance := calculateVariance(data, mean)
	fmt.Printf("Point estimate of the population variance: %.2f\n", variance)
}
