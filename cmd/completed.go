package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var markTaskAsCompletedCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			fmt.Println("Error: ID is required.")
			return
		}
		markTaskAsCompleted(id)
	},
}

func init() {
	rootCmd.AddCommand(markTaskAsCompletedCmd)
	markTaskAsCompletedCmd.Flags().StringP("id", "i", "", "ID of the task (required)")

}
