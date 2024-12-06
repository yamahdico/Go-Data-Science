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

// tValue returns the t-distribution value for given confidence level and degrees of freedom
func tValue(degreesOfFreedom int) float64 {
	// Using a lookup for simplicity, for 95% confidence level and small sample size
	tDistribution := map[int]float64{
		1: 12.71,
		2: 4.303,
		3: 3.182,
		4: 2.776,
		5: 2.571,
	}
	return tDistribution[degreesOfFreedom]
}

func main() {
	// Sample data
	data := []float64{70, 75, 80, 85, 90}

	// Calculate the sample mean
	mean := calculateMean(data)

	// Calculate the sample variance
	variance := calculateVariance(data, mean)
	stdDev := math.Sqrt(variance)

	// Sample size
	n := len(data)

	// Degrees of freedom
	df := n - 1

	// t-distribution value for 95% confidence level
	t := tValue(df)

	// Calculate the margin of error
	marginOfError := t * (stdDev / math.Sqrt(float64(n)))

	// Calculate the confidence interval
	lowerBound := mean - marginOfError
	upperBound := mean + marginOfError

	fmt.Printf("95%% Confidence Interval for the population mean: [%.2f, %.2f]\n", lowerBound, upperBound)
}
