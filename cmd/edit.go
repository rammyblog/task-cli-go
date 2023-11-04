package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a task",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		duration, _ := cmd.Flags().GetString("duration")
		id, _ := cmd.Flags().GetString("id")

		if id == "" {
			fmt.Println("Error: ID is required.")
			return
		}

		editTaskInDb(id, title, duration)
		fmt.Printf("Edited task: %s\n", title)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringP("title", "t", "", "Title of the task (required)")
	editCmd.Flags().StringP("duration", "d", "", "Deadline of the task (format: '5m, 2h30m') (required)")
	editCmd.Flags().StringP("id", "i", "", "ID of the task (required)")

}
