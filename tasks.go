package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int
	Description string
	TotalTime   time.Duration
	StartTime   time.Time
}

type TaskData struct {
	Tasks         []Task
	CurrentTaskID int
}

var taskData TaskData

func resetTasks(force bool) {
	if force {
		// Delete all tasks
		taskData.Tasks = []Task{}
		fmt.Println("All tasks have been deleted.")
	} else {
		// Reset times for all tasks
		for i := range taskData.Tasks {
			taskData.Tasks[i].TotalTime = 0
			taskData.Tasks[i].StartTime = time.Time{}
		}
		fmt.Println("Task times have been reset.")
	}

	saveTaskData()
}

func startTask(taskID int) {
	if taskData.CurrentTaskID != 0 {
		stopTask()
	}

	for i := range taskData.Tasks {
		if taskData.Tasks[i].ID == taskID {
			taskData.Tasks[i].StartTime = time.Now()
			taskData.CurrentTaskID = taskID
			fmt.Printf("Started timer for task: %s\n", taskData.Tasks[i].Description)
			saveTaskData()
			return
		}
	}

	fmt.Printf("Task with ID %d not found\n", taskID)
}

func stopTask() {
	if taskData.CurrentTaskID == 0 {
		fmt.Println("No task is currently running")
		return
	}

	for i := range taskData.Tasks {
		if taskData.Tasks[i].ID == taskData.CurrentTaskID {
			duration := time.Since(taskData.Tasks[i].StartTime)
			taskData.Tasks[i].TotalTime += duration

			fmt.Printf("Stopped timer for task: %s\n", taskData.Tasks[i].Description)
			fmt.Printf("Time spent: %s\n", duration.Round(time.Second))

			taskData.Tasks[i].StartTime = time.Time{}
			taskData.CurrentTaskID = 0

			saveTaskData()
			return
		}
	}

	fmt.Println("Error: Current task not found in the task list")
}

func addTask(name string) {
	task := Task{ID: len(taskData.Tasks) + 1, Description: name}
	taskData.Tasks = append(taskData.Tasks, task)

	saveTaskData()

	fmt.Printf("Task added: %s\n", task.Description)
}

func renameTask(taskID int, newName string) {
	for i := range taskData.Tasks {
		if taskData.Tasks[i].ID == taskID {
			oldName := taskData.Tasks[i].Description
			taskData.Tasks[i].Description = newName
			fmt.Printf("Task renamed from '%s' to '%s'\n", oldName, newName)
			saveTaskData()
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", taskID)
}

func listTaskData() {
	for _, task := range taskData.Tasks {
		totalTime := task.TotalTime

		// Add on current time
		if task.ID == taskData.CurrentTaskID {
			totalTime += time.Since(task.StartTime)
		}

		// Convert the duration to hours and minutes
		hours := int(totalTime.Hours())
		minutes := int(totalTime.Minutes()) % 60

		// Format the time string
		timeStr := fmt.Sprintf("%dh %dm", hours, minutes)

		fmt.Printf("%d: %s (Total time: %s)\n", task.ID, task.Description, timeStr)
	}
}

func printCurrentTask() {
	if taskData.CurrentTaskID == 0 {
		fmt.Println("No task is currently running")
		return
	}
	for _, task := range taskData.Tasks {
		if task.ID == taskData.CurrentTaskID {
			duration := time.Since(task.StartTime)
			fmt.Printf("Current task: %s (ID: %d)\n", task.Description, task.ID)
			fmt.Printf("Running for: %s\n", duration.Round(time.Second))
			return
		}
	}
	fmt.Println("Error: Current task not found in the task list")
}

func deleteTask(taskID int) {
	for i, task := range taskData.Tasks {
		if task.ID == taskID {
			taskData.Tasks = append(taskData.Tasks[:i], taskData.Tasks[i+1:]...)
			fmt.Printf("Task with ID %d has been deleted\n", taskID)
			regenerateTaskIds()
			saveTaskData()
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", taskID)
}

func regenerateTaskIds() {
	for i := range taskData.Tasks {
		taskData.Tasks[i].ID = i + 1
	}
}
