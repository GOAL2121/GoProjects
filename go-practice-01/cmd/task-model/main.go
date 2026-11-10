package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t Task) Summary() string {
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

func (t *Task) Complete() {
	t.Status = "done"
}

func (t *Task) SetPriority(priority int) bool {
	if priority >= 1 && priority <= 3 {
		t.Priority = priority
		return true
	}
	return false
}

func findTaskByID(tasks []Task, id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

func tasksByStatus(tasks []Task, status string) []Task {
	task_status := []Task{}
	for _, task := range tasks {
		if task.Status == status {
			task_status = append(task_status, task)
		}
	}
	return task_status
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}

	found := findTaskByID(tasks, 2)
	if found != nil {
		found.SetPriority(3)
		found.Complete()
	}

	count_done := 0
	for _, task := range tasks {
		if task.Status == "done" {
			count_done++
		}
	}
	fmt.Printf("Задач завершено: %d\n", count_done)

	for _, task := range tasks {
		fmt.Println(task.Summary())
	}
	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}

	fmt.Println(tasksByStatus(tasks, "done"))
}
