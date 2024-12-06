package main

import (
	"fmt"
	"math"
)

// calculateStandardErrorForProportion محاسبه خطای استاندارد نسبت برای یک نمونه
func calculateStandardErrorForProportion(p float64, sampleSize int) float64 {
	if sampleSize <= 0 || p < 0 || p > 1 {
		return 0
	}
	return math.Sqrt((p * (1 - p)) / float64(sampleSize))
}

func main() {
	// نسبت جامعه
	p := 0.6

	// اندازه نمونه
	sampleSize := 100

	// محاسبه خطای استاندارد نسبت
	standardError := calculateStandardErrorForProportion(p, sampleSize)

	fmt.Printf("Standard Error of the Proportion: %.4f\n", standardError)
}
