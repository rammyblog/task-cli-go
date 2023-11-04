package cmd

import (
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "view",
	Short: "Read all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		loadTasksFromFile()
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
