package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gt",
	Short: "GoToDo is a text based todo app",
	Long:  `A Fast daily todo cli app with built in reminders`,
	Run: func(cmd *cobra.Command, args []string) {
		readTodosFromFile()
		refreshTodoDueDatesIfNeeded()
		printTodos(false)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
