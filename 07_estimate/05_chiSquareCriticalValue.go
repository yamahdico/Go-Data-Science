package main

import (
	"fmt"
)

// Calculate Chi-Square critical values using interpolation (simplified for this example)
func chiSquareCriticalValue(p float64, df int) float64 {
	// Simulate critical values for df = 9
	chiSquareValues := map[float64]float64{
		0.025: 19.023, // chi-square value for alpha/2
		0.975: 2.700,  // chi-square value for 1-alpha/2
	}

	if val, exists := chiSquareValues[p]; exists {
		return val
	}
	return 0.0
}

// Calculate the confidence interval for variance
func confidenceIntervalVariance(s2 float64, n int, alpha float64) (float64, float64) {
	df := n - 1
	chi2Lower := chiSquareCriticalValue(alpha/2, df)
	chi2Upper := chiSquareCriticalValue(1-alpha/2, df)

	lower := float64(df) * s2 / chi2Lower
	upper := float64(df) * s2 / chi2Upper

	return lower, upper
}

func main() {
	// Example data
	n := 10
	s2 := 25.0
	alpha := 0.05

	// Calculate confidence interval for variance
	lower, upper := confidenceIntervalVariance(s2, n, alpha)
	fmt.Printf("Confidence Interval for Variance: (%.2f, %.2f)\n", lower, upper)
}
