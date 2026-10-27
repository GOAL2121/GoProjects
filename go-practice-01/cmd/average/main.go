package main

import "fmt"

func main() {
	var min1 int = 1
	var min2 int = 1
	var min3 int = 2
	var sum int = min1 + min2 + min3
	var sr_sum float64 = float64(sum) / 3

	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Среднее арифмитическое: %.2f\n", sr_sum)

}
