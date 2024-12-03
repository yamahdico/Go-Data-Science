package main

import (
	"fmt"
)

func main() {
	// تعریف مجموعه‌ها
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	// اضافه کردن مقادیر به مجموعه‌ها
	set1[3] = true
	set1[7] = true
	set1[5] = true

	set2[5] = true
	set2[8] = true

	fmt.Println("مجموعه اول:", set1)
	fmt.Println("مجموعه دوم:", set2)

	// حذف یک عنصر از مجموعه
	delete(set1, 3)
	fmt.Println("مجموعه اول پس از حذف عنصر 3:", set1)

	// بررسی وجود عنصر
	if set1[7] {
		fmt.Println("عنصر 7 در مجموعه اول وجود دارد")
	} else {
		fmt.Println("عنصر 7 در مجموعه اول وجود ندارد")
	}

	// عملیات اجتماع
	union := make(map[int]bool)
	for k := range set1 {
		union[k] = true
	}
	for k := range set2 {
		union[k] = true
	}
	fmt.Println("اجتماع مجموعه‌ها:", union)

	// عملیات اشتراک
	intersection := make(map[int]bool)
	for k := range set1 {
		if set2[k] {
			intersection[k] = true
		}
	}
	fmt.Println("اشتراک مجموعه‌ها:", intersection)

	// عملیات تفاضل (set1 - set2)
	difference := make(map[int]bool)
	for k := range set1 {
		if !set2[k] {
			difference[k] = true
		}
	}
	fmt.Println("تفاضل (set1 - set2):", difference)

	// عملیات تفاضل متقارن
	symmetricDifference := make(map[int]bool)
	for k := range set1 {
		if !set2[k] {
			symmetricDifference[k] = true
		}
	}
	for k := range set2 {
		if !set1[k] {
			symmetricDifference[k] = true
		}
	}
	fmt.Println("تفاضل متقارن مجموعه‌ها:", symmetricDifference)
}
