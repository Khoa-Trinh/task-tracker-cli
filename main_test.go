// main_test.go
package main

import (
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	var tasks []Task
	desc := "Test add"
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
	if len(tasks) != 1 || tasks[0].Description != desc {
		t.Errorf("AddTask failed")
	}
}

func TestUpdateTask(t *testing.T) {
	tasks := []Task{{ID: 1, Description: "Old", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}}
	i, task := findTask(tasks, 1)
	if task == nil {
		t.Fatal("Task not found")
	}
	task.Description = "Updated"
	task.UpdatedAt = time.Now()
	tasks[i] = *task
	if tasks[0].Description != "Updated" {
		t.Errorf("UpdateTask failed")
	}
}

func TestDeleteTask(t *testing.T) {
	tasks := []Task{{ID: 1, Description: "Delete", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}}
	i, _ := findTask(tasks, 1)
	tasks = append(tasks[:i], tasks[i+1:]...)
	if len(tasks) != 0 {
		t.Errorf("DeleteTask failed")
	}
}

func TestMarkInProgress(t *testing.T) {
	tasks := []Task{{ID: 1, Description: "Progress", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}}
	i, task := findTask(tasks, 1)
	task.Status = "in-progress"
	task.UpdatedAt = time.Now()
	tasks[i] = *task
	if tasks[0].Status != "in-progress" {
		t.Errorf("MarkInProgress failed")
	}
}

func TestMarkDone(t *testing.T) {
	tasks := []Task{{ID: 1, Description: "Done", Status: "in-progress", CreatedAt: time.Now(), UpdatedAt: time.Now()}}
	i, task := findTask(tasks, 1)
	task.Status = "done"
	task.UpdatedAt = time.Now()
	tasks[i] = *task
	if tasks[0].Status != "done" {
		t.Errorf("MarkDone failed")
	}
}

func TestListTasks(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "A", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Description: "B", Status: "done", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	count := 0
	for _, tsk := range tasks {
		if tsk.Status == "done" {
			count++
		}
	}
	if count != 1 {
		t.Errorf("ListTasks filter failed")
	}
}
