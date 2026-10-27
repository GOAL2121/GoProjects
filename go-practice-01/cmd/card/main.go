package main

import "fmt"

func main() {
	var member string = "Полина"
	var group string = "Учебная-01"
	const course = "Бэкенд-разработка на языке Go"
	finished_tasks := 5
	var interest bool = true

	fmt.Println("Участник: ", member)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Printf("Завершено работ: %d\n", finished_tasks)
	fmt.Printf("Интерес к бэкенду: %t\n", interest)

}
