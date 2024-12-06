package main

import (
	"fmt"
	"math"
)

// تابع محاسبه فاکتوریل
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// توزیع دوجمله‌ای نمونه‌ای: محاسبه احتمال k موفقیت در n آزمایش با احتمال موفقیت p
func binomialSampling(n, k int, p float64) float64 {
	combination := float64(factorial(n)) / (float64(factorial(k)) * float64(factorial(n-k)))
	return combination * math.Pow(p, float64(k)) * math.Pow(1-p, float64(n-k))
}

// توزیع نمونه‌ای نرمال: محاسبه چگالی احتمال در نقطه x برای میانگین μ و واریانس σ^2/n
func normalSampling(x, mu, sigma float64, n int) float64 {
	variance := sigma * sigma / float64(n)
	return (1 / (math.Sqrt(2 * math.Pi * variance))) * math.Exp(-math.Pow(x-mu, 2)/(2*variance))
}

// توزیع t نمونه‌ای: محاسبه احتمال t
func tDistribution(x, mu, s float64, n int) float64 {
	df := float64(n - 1)
	t := (x - mu) / (s / math.Sqrt(float64(n)))
	return (1 / (math.Sqrt(df * math.Pi))) * math.Exp(-math.Pow(t, 2)/(2*df))
}

func main() {
	// نمونه‌ای از توزیع دوجمله‌ای
	n, k := 10, 5
	p := 0.5
	fmt.Printf("Binomial Sampling Probability: P(X=%d) = %.5f\n", k, binomialSampling(n, k, p))

	// نمونه‌ای از توزیع نرمال
	x, mu, sigma := 0.0, 0.0, 1.0
	nSample := 30
	fmt.Printf("Normal Sampling Probability Density: f(%.2f) = %.5f\n", x, normalSampling(x, mu, sigma, nSample))

	// نمونه‌ای از توزیع t
	x, mu, s := 1.5, 0.0, 1.0
	nSample = 30
	fmt.Printf("t Distribution Probability: t = %.5f\n", tDistribution(x, mu, s, nSample))
}
