package cmd

import (
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "view-all",
	Short: "Read all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		readAllTasks()
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
