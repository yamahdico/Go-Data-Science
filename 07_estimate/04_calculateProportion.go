package main

import (
	"fmt"
	"math"
)

// calculateProportion محاسبه نسبت نمونه
func calculateProportion(successes, sampleSize int) float64 {
	return float64(successes) / float64(sampleSize)
}

// zValue بازگرداندن مقدار z برای سطح اطمینان 95%
func zValue() float64 {
	// Z-value for 95% confidence level is approximately 1.96
	return 1.96
}

// calculateConfidenceInterval محاسبه فاصله اطمینان
func calculateConfidenceInterval(pHat float64, n int, z float64) (float64, float64) {
	// محاسبه انحراف معیار
	standardError := math.Sqrt(pHat * (1 - pHat) / float64(n))

	// محاسبه حاشیه خطا
	marginOfError := z * standardError

	// محاسبه حدود بالایی و پایینی فاصله اطمینان
	lowerBound := pHat - marginOfError
	upperBound := pHat + marginOfError

	return lowerBound, upperBound
}

func main() {
	// داده‌های نمونه
	successes := 120
	sampleSize := 200

	// محاسبه نسبت نمونه
	pHat := calculateProportion(successes, sampleSize)

	// مقدار z برای سطح اطمینان 95%
	z := zValue()

	// محاسبه فاصله اطمینان
	lowerBound, upperBound := calculateConfidenceInterval(pHat, sampleSize, z)

	fmt.Printf("95%% Confidence Interval for the population proportion: [%.4f, %.4f]\n", lowerBound, upperBound)
}
