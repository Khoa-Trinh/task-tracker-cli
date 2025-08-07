package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

const fileName = "tasks.json"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func loadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.Open(fileName)
	if errors.Is(err, os.ErrNotExist) {
		return tasks, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", " ")
	return enc.Encode(tasks)
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func findTask(tasks []Task, id int) (int, *Task) {
	for i, t := range tasks {
		if t.ID == id {
			return i, &tasks[i]
		}
	}
	return -1, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}
	command := os.Args[1]
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <description>")
			return
		}
		desc := os.Args[2]
		id := nextID(tasks)
		now := time.Now()
		task := Task{
			ID:          id,
			Description: desc,
			Status:      "todo",
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		tasks = append(tasks, task)
		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Printf("Task added successfully (ID: %d)\n", id)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> <status>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		desc := os.Args[3]
		i, task := findTask(tasks, id)
		if task == nil {
			fmt.Printf("Task with ID %d not found\n", id)
			return
		}
		task.Description = desc
		task.UpdatedAt = time.Now()
		tasks[i] = *task
		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Printf("Task updated successfully (ID: %d)\n", id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		i, _ := findTask(tasks, id)
		if i == -1 {
			fmt.Printf("Task with ID %d not found\n", id)
			return
		}
		tasks = append(tasks[:i], tasks[i+1:]...)
		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Printf("Task deleted successfully (ID: %d)\n", id)

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		i, task := findTask(tasks, id)
		if task == nil {
			fmt.Printf("Task with ID %d not found\n", id)
			return
		}
		task.Status = "in-progress"
		task.UpdatedAt = time.Now()
		tasks[i] = *task
		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Printf("Task marked as in-progress (ID: %d)\n", id)

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-done <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		i, task := findTask(tasks, id)
		if task == nil {
			fmt.Printf("Task with ID %d not found\n", id)
			return
		}
		task.Status = "done"
		task.UpdatedAt = time.Now()
		tasks[i] = *task
		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Printf("Task marked as done (ID: %d)\n", id)

	case "list":
		filter := ""
		if len(os.Args) > 2 {
			filter = os.Args[2]
		}
		for _, t := range tasks {
			if filter == "" || t.Status == filter {
				fmt.Printf("ID: %d | %s | %s | Created: %s | Updated: %s\n",
					t.ID, t.Description, t.Status, t.CreatedAt.Format(time.RFC3339), t.UpdatedAt.Format(time.RFC3339))
			}
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
