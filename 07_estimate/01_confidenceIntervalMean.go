package main

import (
	"fmt"
	"math"
)

func mean(data []float64) float64 {
	var sum float64
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

func variance(data []float64) float64 {
	m := mean(data)
	var sum float64
	for _, value := range data {
		sum += (value - m) * (value - m)
	}
	return sum / float64(len(data)-1)
}

func confidenceIntervalMean(data []float64, confidenceLevel float64) (float64, float64) {
	m := mean(data)
	s := math.Sqrt(variance(data))
	n := float64(len(data))
	z := 1.96 // مقدار z برای 95% فاصله اطمینان
	marginOfError := z * s / math.Sqrt(n)
	return m - marginOfError, m + marginOfError
}

func main() {
	data := []float64{10.0, 12.5, 9.7, 11.3, 13.1, 10.9}
	populationMeanEstimate := mean(data)
	populationVarianceEstimate := variance(data)
	lower, upper := confidenceIntervalMean(data, 0.95)

	fmt.Printf("تخمین میانگین جمعیت: %.2f\n", populationMeanEstimate)
	fmt.Printf("تخمین واریانس جمعیت: %.2f\n", populationVarianceEstimate)
	fmt.Printf("فاصله اطمینان 95%% برای میانگین: (%.2f, %.2f)\n", lower, upper)
}
