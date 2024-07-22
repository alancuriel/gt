package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeCmd)
}

var completeCmd = &cobra.Command{
	Use:   "c",
	Short: "Complete tasks todo",
	Long:  `Mark a task/todo as a complete`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a todo number to complete eg. 1")
			return
		}

		todoToRemove, err := strconv.ParseInt(args[0], 10, 64)

		if err != nil {
			fmt.Println("Please provide a valid todo number to complete eg. 1")
			return
		}

		readTodosFromFile()
		refreshTodoDueDatesIfNeeded()

		var count int64 = 0
		completed := false

		for i,todo := range todos {
			if !todo.completed {
				if count+1 == todoToRemove {
					todos[i].completed = true
					completed = true
					fmt.Printf("Todo #%d, marked as complete\n\n", todoToRemove)
					break
				}
				count += 1
			}
		}

		if !completed {
			fmt.Printf("Could not find todo #%d to remove.\n\n", todoToRemove)
			printTodos(false)
			return
		}

		printTodos(false)

		writeTodosToFile()
	},
}
