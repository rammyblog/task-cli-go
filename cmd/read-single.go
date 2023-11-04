package cmd

import (
	"github.com/spf13/cobra"
)

var readSingleCmd = &cobra.Command{
	Use:   "view",
	Short: "Read all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		readSingleTask(id)
	},
}

func init() {
	rootCmd.AddCommand(readSingleCmd)
}