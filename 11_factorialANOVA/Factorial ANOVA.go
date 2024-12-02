package main

import (
	"fmt"
	"math"
)

// Sample data: each slice represents a group based on two factors (e.g., training type and gender)
var data = [][]float64{
	{85, 90, 88, 86}, // Group 1 (Training A, Male)
	{78, 82, 80, 79}, // Group 2 (Training A, Female)
	{92, 91, 94, 93}, // Group 3 (Training B, Male)
	{84, 85, 83, 82}, // Group 4 (Training B, Female)
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
