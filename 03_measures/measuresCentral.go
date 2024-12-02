package main

// Measures of Central Tendency

import (
	"fmt"
	"sort"
)

// Function to calculate the mean of the data set
func mean(data []float64) float64 {
	sum := 0.0
	// Calculate the sum of all values
	for _, value := range data {
		sum += value
	}
	// Divide the sum by the number of values to get the mean
	return sum / float64(len(data))
}

// Function to calculate the median of the data set
func median(data []float64) float64 {
	// Sort the data
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		// If the number of data points is even, return the average of the two middle values
		return (data[n/2-1] + data[n/2]) / 2
	}
	// If the number of data points is odd, return the middle value
	return data[n/2]
}

// Function to calculate the mode of the data set
func mode(data []float64) float64 {
	frequency := make(map[float64]int)
	maxCount := 0
	var mode float64

	// Count the frequency of each value
	for _, value := range data {
		frequency[value]++
		if frequency[value] > maxCount {
			maxCount = frequency[value]
			mode = value
		}
	}

	return mode
}

func main() {
	// Define a sample data set
	data := []float64{4, 8, 6, 5, 3, 9, 4}

	// Calculate mean, median, and mode
	m := mean(data)
	md := median(data)
	mdm := mode(data)

	// Print the results
	fmt.Printf("Mean: %.2f\n", m)    // Mean of the data set
	fmt.Printf("Median: %.2f\n", md) // Median of the data set
	fmt.Printf("Mode: %.2f\n", mdm)  // Mode of the data set
}
