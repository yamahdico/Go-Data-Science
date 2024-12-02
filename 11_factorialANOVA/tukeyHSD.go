package main

import (
	"fmt"
	"math"
)

// داده‌های نمونه
var group1 = []float64{85, 88, 90}
var group2 = []float64{80, 82, 83}
var group3 = []float64{88, 90, 92}

// تابع محاسبه میانگین
func mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// تابع محاسبه واریانس
func variance(data []float64) float64 {
	m := mean(data)
	variance := 0.0
	for _, value := range data {
		variance += math.Pow(value-m, 2)
	}
	return variance / float64(len(data)-1)
}

// تابع محاسبه فاصله توکی بین دو گروه
func tukeyHSD(group1, group2 []float64, mse float64, n float64) float64 {
	meanDiff := math.Abs(mean(group1) - mean(group2))
	se := math.Sqrt(mse * 2 / n)
	return meanDiff / se
}

func main() {
	// محاسبه میانگین و واریانس برای هر گروه
	mean1 := mean(group1)
	mean2 := mean(group2)
	mean3 := mean(group3)

	// محاسبه واریانس درون‌گروهی (MSE)
	var pooledVariance float64
	pooledVariance += variance(group1)
	pooledVariance += variance(group2)
	pooledVariance += variance(group3)
	mse := pooledVariance / float64(len(group1)+len(group2)+len(group3)-3)

	// تعداد مشاهدات در هر گروه (فرض می‌کنیم هر گروه یک تعداد دارد)
	n := float64(len(group1))

	// محاسبه فاصله‌های توکی
	tukey12 := tukeyHSD(group1, group2, mse, n)
	tukey13 := tukeyHSD(group1, group3, mse, n)
	tukey23 := tukeyHSD(group2, group3, mse, n)

	fmt.Printf("Tukey HSD between group 1 and 2: %.2f\n", tukey12)
	fmt.Printf("Tukey HSD between group 1 and 3: %.2f\n", tukey13)
	fmt.Printf("Tukey HSD between group 2 and 3: %.2f\n", tukey23)

	// تفسیر نتایج: اگر فاصله توکی بزرگتر از مقدار بحرانی باشد، تفاوت معنادار است.
}
