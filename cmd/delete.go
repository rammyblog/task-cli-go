package cmd

import (
	"github.com/spf13/cobra"
)

var deleteSingleCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		deleteTask(id)
	},
}

func init() {
	rootCmd.AddCommand(deleteSingleCmd)
}
