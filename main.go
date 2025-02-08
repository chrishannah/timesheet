package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Task struct {
    ID          int
    Description string
    TotalTime   time.Duration
    StartTime   time.Time
}

var tasks []Task
var currentTask *Task

func main() {
	loadTasks()

	rootCmd := &cobra.Command{Use: "timesheet"}
	rootCmd.AddCommand(addCmd, listCmd, resetCmd)
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

		task := Task{ID: len(tasks) + 1, Description: args[0]}
		tasks = append(tasks, task)

		saveTasks()

		fmt.Printf("Task added: %s\n", task.Description)
	},
}

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		for _, task := range tasks {
			fmt.Printf("%d: %s (Total time: %s)\n", task.ID, task.Description, task.TotalTime)
		}
	},
}

var resetCmd = &cobra.Command{
	Use: "reset",
	Short: "Reset all tasks.",
	Long: "Reset the timing of all tasks, or use the --force flag to also delete all tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		resetTasks(force)
	},
}

func saveTasks() {
    file, _ := json.MarshalIndent(tasks, "", " ")
    _ = os.WriteFile("tasks.json", file, 0644)
}

func loadTasks() {
    file, _ := os.ReadFile("tasks.json")
    _ = json.Unmarshal(file, &tasks)
}

func resetTasks(force bool) {
	if force {
        // Delete all tasks
        tasks = []Task{}
        fmt.Println("All tasks have been deleted.")
    } else {
        // Reset times for all tasks
        for i := range tasks {
            tasks[i].TotalTime = 0
            tasks[i].StartTime = time.Time{}
        }
        fmt.Println("Task times have been reset.")
    }

    saveTasks()
}
