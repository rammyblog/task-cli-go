package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteSingleCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			fmt.Println("Error: ID is required.")
			return
		}
		deleteTask(id)
	},
}

func init() {
	rootCmd.AddCommand(deleteSingleCmd)
	deleteSingleCmd.Flags().StringP("id", "id", "", "ID of the task (required)")

}
