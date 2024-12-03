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

// توزیع دوجمله‌ای: محاسبه احتمال k موفقیت در n آزمایش با احتمال موفقیت p
func binomial(n, k int, p float64) float64 {
	combination := float64(factorial(n)) / (float64(factorial(k)) * float64(factorial(n-k)))
	return combination * math.Pow(p, float64(k)) * math.Pow(1-p, float64(n-k))
}

// توزیع پواسون: محاسبه احتمال k رخداد با نرخ λ
func poisson(k int, lambda float64) float64 {
	return math.Pow(lambda, float64(k)) * math.Exp(-lambda) / float64(factorial(k))
}

// توزیع نرمال: محاسبه احتمال در نقطه x برای میانگین μ و انحراف معیار σ
func normal(x, mu, sigma float64) float64 {
	return (1 / (sigma * math.Sqrt(2*math.Pi))) * math.Exp(-math.Pow(x-mu, 2)/(2*math.Pow(sigma, 2)))
}

func main() {
	// نمونه‌ای از استفاده از توزیع دوجمله‌ای
	n, k := 10, 5
	p := 0.5
	fmt.Printf("Binomial Probability: P(X=%d) = %.5f\n", k, binomial(n, k, p))

	// نمونه‌ای از استفاده از توزیع پواسون
	lambda := 3.0
	k = 2
	fmt.Printf("Poisson Probability: P(X=%d) = %.5f\n", k, poisson(k, lambda))

	// نمونه‌ای از استفاده از توزیع نرمال
	x := 0.0
	mu, sigma := 0.0, 1.0
	fmt.Printf("Normal Probability Density: f(%.2f) = %.5f\n", x, normal(x, mu, sigma))
}
