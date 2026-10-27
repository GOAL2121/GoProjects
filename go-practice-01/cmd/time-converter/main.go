package main

import "fmt"

func main() {
	var totalMinutes int
	fmt.Print("Введите количество минут: ")

	_, err := fmt.Scan(&totalMinutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if totalMinutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}

	const min_h = 60
	var hours int = totalMinutes / min_h
	var min int = totalMinutes % min_h

	fmt.Printf("Полных часов: %d\n", hours)
	fmt.Printf("Оставшихся минут: %d\n", min)
}
