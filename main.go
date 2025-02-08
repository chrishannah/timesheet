package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
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

func main() {
	loadTaskData()

	rootCmd := &cobra.Command{Use: "timesheet"}
	rootCmd.AddCommand(addCmd, listCmd, resetCmd, startCmd, stopCmd, currentCmd)
	rootCmd.Execute()
}

func init() {
    // Add the force flag to the reset command
    resetCmd.Flags().BoolP("force", "f", false, "Delete all tasks instead of just resetting times")
}

var addCmd = &cobra.Command{
	Use: "add [description]",
	Short: "Add new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a task name")
			return
		}

		task := Task{ID: len(taskData.Tasks) + 1, Description: args[0]}
		taskData.Tasks = append(taskData.Tasks, task)

		saveTaskData()

		fmt.Printf("Task added: %s\n", task.Description)
	},
}

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		for _, task := range taskData.Tasks {
			// Convert the duration to hours and minutes
            hours := int(task.TotalTime.Hours())
            minutes := int(task.TotalTime.Minutes()) % 60

            // Format the time string
            timeStr := fmt.Sprintf("%dh %dm", hours, minutes)

			fmt.Printf("%d: %s (Total time: %s)\n", task.ID, task.Description, timeStr)
		}
	},
}

var resetCmd = &cobra.Command{
	Use: "reset",
	Short: "Reset all tasks.",
	Long: "Reset the timing of all tasks, or use the -	orce flag to also delete all tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		resetTasks(force)
	},
}

var startCmd = &cobra.Command{
    Use:   "start [id]",
    Short: "Start timer for a task",
    Long:  `Start the timer for a task with the given ID. If a timer is already running for another task, it will be stopped first.`,
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 1 {
            fmt.Println("Please provide a task ID")
            return
        }

        taskID, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Println("Invalid task ID. Please provide a number.")
            return
        }

        startTask(taskID)
    },
}

var stopCmd = &cobra.Command{
    Use:   "stop",
    Short: "Stop current timer",
    Long:  "If a timer is running, stop it.",
    Run: func(cmd *cobra.Command, args []string) {
        stopTask()
    },
}

var currentCmd = &cobra.Command{
    Use:   "current",
    Short: "Show the currently running task",
    Run: func(cmd *cobra.Command, args []string) {
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
    },
}

func saveTaskData() {
    file, _ := json.MarshalIndent(taskData, "", " ")
    _ = os.WriteFile("taskdata.json", file, 0644)
}

func loadTaskData() {
    file, err := os.ReadFile("taskdata.json")
    if err != nil {
        // If the file doesn't exist, initialize with empty data
        taskData = TaskData{Tasks: []Task{}}
        return
    }
    _ = json.Unmarshal(file, &taskData)
}

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
