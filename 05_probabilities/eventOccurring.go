package main

import "fmt"

// تابعی برای محاسبه احتمال
func probability(successfulOutcomes, totalOutcomes int) float64 {
	// جلوگیری از تقسیم بر صفر
	if totalOutcomes == 0 {
		return 0
	}
	// محاسبه احتمال
	return float64(successfulOutcomes) / float64(totalOutcomes)
}

func main() {
	// تعداد حالات مطلوب و کل حالات ممکن برای مثال‌های مختلف
	success := 1 // مثلا برای پرتاب تاس و به دست آوردن عدد 4
	total := 6   // چون تاس شش وجه دارد

	// محاسبه احتمال
	p := probability(success, total)

	// نمایش احتمال
	fmt.Printf("احتمال به دست آوردن عدد 4 در پرتاب یک تاس شش وجهی: %.2f\n", p)
}
