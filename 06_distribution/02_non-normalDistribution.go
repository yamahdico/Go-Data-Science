package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// تابع تولید نمونه تصادفی از یک توزیع غیر نرمال (مثلاً توزیع یکنواخت)
func generateUniformSample(n int) []float64 {
	sample := make([]float64, n)
	for i := 0; i < n; i++ {
		sample[i] = rand.Float64() // تولید عدد تصادفی بین 0 و 1
	}
	return sample
}

// تابع محاسبه میانگین یک نمونه
func calculateMean(sample []float64) float64 {
	var sum float64
	for _, value := range sample {
		sum += value
	}
	return sum / float64(len(sample))
}

// تابع محاسبه توزیع نرمال برای نمایش و مقایسه
func normalDistribution(x, mean, stddev float64) float64 {
	return (1 / (stddev * math.Sqrt(2*math.Pi))) * math.Exp(-math.Pow(x-mean, 2)/(2*stddev*stddev))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// اندازه نمونه
	n := 30
	numSamples := 1000

	var means []float64

	// تولید میانگین‌های نمونه‌های مختلف
	for i := 0; i < numSamples; i++ {
		sample := generateUniformSample(n)
		mean := calculateMean(sample)
		means = append(means, mean)
	}

	// محاسبه میانگین و انحراف معیار توزیع نمونه‌ای میانگین‌ها
	var meanSum, stddevSum float64
	for _, mean := range means {
		meanSum += mean
	}
	meanOfMeans := meanSum / float64(numSamples)

	for _, mean := range means {
		stddevSum += math.Pow(mean-meanOfMeans, 2)
	}
	stddevOfMeans := math.Sqrt(stddevSum / float64(numSamples))

	fmt.Printf("Mean of Means: %.4f\n", meanOfMeans)
	fmt.Printf("Standard Deviation of Means: %.4f\n", stddevOfMeans)

	// نمایش توزیع نرمال برای مقایسه
	fmt.Println("Normal Distribution Probability Density for Mean of Means:")
	for i := -3; i <= 3; i++ {
		x := float64(i) * stddevOfMeans / 3
		fmt.Printf("f(%.2f) = %.5f\n", x, normalDistribution(x, meanOfMeans, stddevOfMeans))
	}
}
