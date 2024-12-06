package main

import (
	"fmt"
	"math"
)

// calculateStandardError محاسبه خطای استاندارد میانگین برای یک نمونه
func calculateStandardError(stdDev float64, sampleSize int) float64 {
	if sampleSize <= 0 {
		return 0
	}
	return stdDev / math.Sqrt(float64(sampleSize))
}

func main() {
	// انحراف معیار جامعه
	stdDev := 10.0

	// اندازه نمونه
	sampleSize := 25

	// محاسبه خطای استاندارد میانگین
	standardError := calculateStandardError(stdDev, sampleSize)

	fmt.Printf("Standard Error of the Mean: %.2f\n", standardError)
}
