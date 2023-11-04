package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		duration := args[1]
		saveTasksToFile(title, duration)
		fmt.Printf("Added task: %s\n", title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
