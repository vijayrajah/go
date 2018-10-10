package main

import (
	"fmt"
)

func sumdiff(a, b int) (int, int) {
	return a + b, a - b
}

func sumavg(nums ...int) (int, float64) {
	total := 0
	var avg float64

	for _, num := range nums {
		total += num
	}
	avg = float64(total) / float64(len(nums))
	return total, avg
}

func main() {
	a := 10
	b := 20
	sum, dif := sumdiff(a, b)
	fmt.Println("Sum is ", sum, " & Difference is ", dif)

	numarr := []int{10, 20, 100, 12}
	arsum, aravg := sumavg(numarr...)
	fmt.Println("Sum is ", arsum, " Avg is ", aravg)
}
