package main

import (
	"fmt"
	"math"
)

// calculateCorrelationCoefficient محاسبه ضریب همبستگی پیرسون
func calculateCorrelationCoefficient(x, y []float64) float64 {
	n := float64(len(x))
	if n == 0 {
		return 0
	}

	// محاسبه مجموع های مورد نیاز
	var sumX, sumY, sumXY, sumXSquare, sumYSquare float64
	for i := 0; i < int(n); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXSquare += x[i] * x[i]
		sumYSquare += y[i] * y[i]
	}

	// محاسبه ضریب همبستگی
	numerator := n*sumXY - sumX*sumY
	denominator := math.Sqrt((n*sumXSquare - sumX*sumX) * (n*sumYSquare - sumY*sumY))

	// بررسی جلوگیری از تقسیم بر صفر
	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}

func main() {
	// داده های نمونه
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 3, 4, 5, 6}

	// محاسبه ضریب همبستگی
	r := calculateCorrelationCoefficient(x, y)

	fmt.Printf("Correlation Coefficient (r): %.2f\n", r)
}
