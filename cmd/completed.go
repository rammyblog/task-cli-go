package cmd

import (
	"github.com/spf13/cobra"
)

var markTaskAsCompletedCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		markTaskAsCompleted(id)
	},
}

func init() {
	rootCmd.AddCommand(markTaskAsCompletedCmd)
}
