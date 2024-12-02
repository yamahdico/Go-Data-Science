package main

import (
	"fmt"
	"math"
)

// Sample data
var data = [][]float64{
	{23, 27, 24, 21},
	{30, 31, 29, 32},
	{22, 21, 23, 24},
	{25, 27, 26, 28},
}

func mean(values []float64) float64 {
	sum := 0.0
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}

// Function to calculate sum of squares total
func sumOfSquaresTotal(data [][]float64, grandMean float64) float64 {
	ssTotal := 0.0
	for _, group := range data {
		for _, value := range group {
			ssTotal += math.Pow(value-grandMean, 2)
		}
	}
	return ssTotal
}

// Main function
func main() {
	// Calculate grand mean
	var allData []float64
	for _, group := range data {
		allData = append(allData, group...)
	}
	grandMean := mean(allData)

	// Calculate sum of squares total
	ssTotal := sumOfSquaresTotal(data, grandMean)
	fmt.Printf("Sum of Squares Total (SS_Total): %.2f\n", ssTotal)

	// Additional calculations for SS_A, SS_B, SS_AB, SS_Error, etc. would be implemented here.
	// The implementation would include calculating means for each level of factors A and B, and then the interactions.

	// For full ANOVA analysis, calculations for degrees of freedom, mean squares, F-statistic and p-values should be added.
}
