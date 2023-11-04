package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		duration, _ := cmd.Flags().GetString("duration")

		if title == "" {
			fmt.Println("Error: Title is required.")
			return
		}

		if duration == "" {
			fmt.Println("Duration is required")
			return
		}
		saveTasksToFile(title, duration)
		fmt.Printf("Added task: %s\n", title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("title", "t", "", "Title of the task (required)")
	addCmd.Flags().StringP("duration", "d", "", "Deadline of the task (format: '5m, 2h30m') (required)")
}
