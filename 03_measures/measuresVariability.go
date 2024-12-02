package main

// Measures of variability

import (
	"fmt"
	"math"
)

// Function to calculate the range of the data set
func calculateRange(data []float64) float64 {
	// Initialize min and max with the first element
	min := data[0]
	max := data[0]

	// Loop through the data to find the minimum and maximum values
	for _, value := range data {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	// Calculate the range by subtracting min from max
	return max - min
}

// Function to calculate the mean (average) of the data set
func mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// Function to calculate the variance of the data set
func variance(data []float64) float64 {
	m := mean(data) // Calculate the mean of the data set
	sum := 0.0
	for _, value := range data {
		// Sum up the squares of the differences between each value and the mean
		sum += math.Pow(value-m, 2)
	}
	// Return the variance (sample variance, using N-1)
	return sum / float64(len(data)-1)
}

// Function to calculate the standard deviation of the data set
func standardDeviation(data []float64) float64 {
	// Standard deviation is the square root of the variance
	return math.Sqrt(variance(data))
}

func main() {
	// Define a sample data set
	data := []float64{4, 8, 6, 5, 3, 9}

	// Calculate range, variance, and standard deviation
	r := calculateRange(data)
	v := variance(data)
	sd := standardDeviation(data)

	// Print the results
	fmt.Printf("Range: %.2f\n", r)
	fmt.Printf("Variance: %.2f\n", v)
	fmt.Printf("Standard Deviation: %.2f\n", sd)
}
