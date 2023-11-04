package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var readSingleCmd = &cobra.Command{
	Use:   "view",
	Short: "Read a single task",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			fmt.Println("Error: ID is required.")
			return
		}
		readSingleTask(id)
	},
}

func init() {
	rootCmd.AddCommand(readSingleCmd)
	readSingleCmd.Flags().StringP("id", "i", "", "ID of the task (required)")

}
