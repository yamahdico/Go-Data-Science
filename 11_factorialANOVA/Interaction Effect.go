package main

//Interaction Effect

import (
	"fmt"
	"math"
)

// داده‌های نمونه: هر اسلایس نمایانگر یک گروه بر اساس دو عامل (مثلاً نوع آموزش و جنسیت)
var data = [][]float64{
	{85, 90, 88, 86}, // گروه 1 (آموزش A، مرد)
	{78, 82, 80, 79}, // گروه 2 (آموزش A، زن)
	{92, 91, 94, 93}, // گروه 3 (آموزش B، مرد)
	{84, 85, 83, 82}, // گروه 4 (آموزش B، زن)
}

func mean(values []float64) float64 {
	sum := 0.0
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}

// تابع برای محاسبه مجموع مربعات تعامل
func sumOfSquaresInteraction(data [][]float64, grandMean float64, rowMeans []float64, colMeans []float64) float64 {
	ssInteraction := 0.0
	for i, row := range data {
		for j, value := range row {
			rowMean := rowMeans[i/2]        // تعداد سطوح عامل A
			colMean := colMeans[j/len(row)] // تعداد سطوح عامل B
			ssInteraction += math.Pow(value-rowMean-colMean+grandMean, 2)
		}
	}
	return ssInteraction
}

func main() {
	// محاسبه میانگین کل
	var allData []float64
	for _, group := range data {
		allData = append(allData, group...)
	}
	grandMean := mean(allData)

	// محاسبه میانگین‌ها برای هر سطح از عامل A و B
	rowMeans := []float64{
		mean(data[0]), // میانگین برای آموزش A
		mean(data[1]), // میانگین برای آموزش B
	}
	colMeans := []float64{
		mean([]float64{data[0][0], data[2][0]}), // میانگین برای مردان
		mean([]float64{data[1][0], data[3][0]}), // میانگین برای زنان
	}

	// محاسبه مجموع مربعات تعامل
	ssInteraction := sumOfSquaresInteraction(data, grandMean, rowMeans, colMeans)
	fmt.Printf("Sum of Squares Interaction (SS_Interaction): %.2f\n", ssInteraction)

	// محاسبات بیشتر برای تحلیل کامل اثرات اصلی و تعاملات می‌تواند در اینجا اضافه شود.
}
