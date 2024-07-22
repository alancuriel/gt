package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "l",
	Short: "List tasks todo",
	Long:  `List tasks todo that are completed`,
	Run: func(cmd *cobra.Command, args []string) {
		readTodosFromFile()
		refreshTodoDueDatesIfNeeded()
		printTodos(false)
	},
}
