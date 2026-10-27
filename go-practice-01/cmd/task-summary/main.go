package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {

	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестный статус"
	}
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	k := 0

	for _, status := range statuses {
		if status == target {
			k++
		}
	}

	return k

}

func findStatus(statuses []string, target string) int {
	for i, status := range statuses {
		if status == target {
			return i
		}
	}
	return -1
}

func completionPercent(statuses []string) float64 {
	if len(statuses) == 0 {
		return 0
	} else {
		return float64(countStatus(statuses, "done")) / float64(len(statuses)) * 100
	}
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	counts := make(map[string]int)

	for _, status := range statuses {
		counts[status]++
	}

	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d\n", counts["new"], counts["in_progress"], counts["done"])

	statuses1 := []string{"new", "in_progress", "done", "done", "done", "delited", "blocked"}
	statuses2 := []string{"done", "done"}
	statuses3 := []string{}

	fmt.Println("Дополнительное задание:")
	fmt.Printf("1) Поиск %d / Доля: %f\n", findStatus(statuses1, "done"), completionPercent(statuses1))
	fmt.Printf("2) Поиск %d / Доля: %f\n", findStatus(statuses2, "done"), completionPercent(statuses2))
	fmt.Printf("3) Поиск %d / Доля: %f\n", findStatus(statuses3, "done"), completionPercent(statuses3))

}
