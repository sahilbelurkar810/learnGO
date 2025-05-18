package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "Todo is a simple CLI task manager",
	}

	var addCmd = &cobra.Command{
		Use:   "add [task description]",
		Short: "Add a new task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			AddTask(args[0])
			SaveTasks()
		},
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			ListTasks()
		},
	}

	var doneCmd = &cobra.Command{
		Use:   "done [task ID]",
		Short: "Mark a task as done",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID")
				return
			}
			MarkDone(id)
			SaveTasks()
		},
	}

	var deleteCmd = &cobra.Command{
		Use:   "delete [task ID]",
		Short: "Delete a task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID")
				return
			}
			DeleteTask(id)
			SaveTasks()
		},
	}

	rootCmd.AddCommand(addCmd, listCmd, doneCmd, deleteCmd)

	// Load tasks first
	err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
