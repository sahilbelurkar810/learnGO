package main

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

var tasks []Task
var nextID = 1

func AddTask(title string) {
	task := Task{
		ID:    nextID,
		Title: title,
		Done:  false,
	}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Added task:", title)
}

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet.")
		return
	}
	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
	}
}
func MarkDone(id int) {
	for i, task := range tasks {
		if task.ID == id {
			if tasks[i].Done {
				fmt.Println("Task already marked as done.")
				return
			}
			tasks[i].Done = true
			fmt.Printf("Marked task %d as done.\n", id)
			return
		}
	}
	fmt.Println("Task not found.")
}

func DeleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Deleted task %d.\n", id)
			return
		}
	}
	fmt.Println("Task not found.")
}
