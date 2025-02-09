package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "timesheet"}

var addCmd = &cobra.Command{
	Use: "add [description]",
	Short: "Add new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a task name")
			return
		}
		addTask(args[0])
	},
}

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		listTaskData()
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
        printCurrentTask()
    },
}

var renameCmd = &cobra.Command{
    Use:   "rename [id] [new name]",
    Short: "Rename a task",
    Long:  "Rename a task by providing its ID and the new name",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 2 {
            fmt.Println("Please provide a task ID and a new name")
            return
        }

        taskID, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Println("Invalid task ID. Please provide a number.")
            return
        }

        newName := args[1]
        renameTask(taskID, newName)
    },
}

var deleteCmd = &cobra.Command{
    Use:   "delete [id]",
    Short: "Delete a task by its ID",
    Long:  "Delete a task from the timesheet by specifying its ID",
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

        deleteTask(taskID)
    },
}

func init() {
    // Add the force flag to the reset command
    resetCmd.Flags().BoolP("force", "f", false, "Delete all tasks instead of just resetting times")

	rootCmd.AddCommand(addCmd, deleteCmd, listCmd, resetCmd, startCmd, stopCmd, currentCmd, renameCmd)

}
