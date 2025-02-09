package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
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
    Use:   "start",
    Short: "Start timer for a task",
    Long:  `Start the timer from a list of already defined tasks.`,
    Run: func(cmd *cobra.Command, args []string) {

        prompt := buildTaskSelectPrompt()
        index, _, err := prompt.Run()

        if err != nil {
            fmt.Printf("Prompt failed %v\n", err)
            return
        }

        task := taskData.Tasks[index]
        startTask(task.ID)
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
    Use:   "rename",
    Short: "Rename a task",
    Long:  "Rename a task by selecting from a list of tasks, and providing a  new name",
    Run: func(cmd *cobra.Command, args []string) {

        select_prompt := buildTaskSelectPrompt()
        index, _, err := select_prompt.Run()

        if err != nil {
            fmt.Printf("Prompt failed %v\n", err)
            return
        }

        name_prompt := promptui.Prompt{
            Label:    "New name",
        }
        name, err := name_prompt.Run()

        if err != nil {
            fmt.Printf("Prompt failed %v\n", err)
            return
        }

        if len(name) < 1 {
            fmt.Println("Please provide a new name for the task")
            return
        }

        task := taskData.Tasks[index]
        renameTask(task.ID, name)
    },
}

var deleteCmd = &cobra.Command{
    Use:   "delete ",
    Short: "Delete a task",
    Long:  "Delete a task from the list of already defined tasks",
    Run: func(cmd *cobra.Command, args []string) {
        prompt := buildTaskSelectPrompt()
        index, _, err := prompt.Run()

        if err != nil {
            fmt.Printf("Prompt failed %v\n", err)
            return
        }

        task := taskData.Tasks[index]
        deleteTask(task.ID)
    },
}

func init() {
    // Add the force flag to the reset command
    resetCmd.Flags().BoolP("force", "f", false, "Delete all tasks instead of just resetting times")

	rootCmd.AddCommand(addCmd, deleteCmd, listCmd, resetCmd, startCmd, stopCmd, currentCmd, renameCmd)

}

func buildTaskSelectPrompt() promptui.Select {
    templates := &promptui.SelectTemplates{
        Label:    "{{ .Description }}",
        Active:   "\U00002605 {{ .Description | cyan }}",
        Inactive: "  {{ .Description | cyan }}",
        Selected: "\U00002605 {{ .Description | red | cyan }}",
    }

    prompt := promptui.Select{
        Label: "Select a task",
        Items: taskData.Tasks,
        Templates: templates,
    }

    return prompt
}
