package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "u",
	Short: "Update todo/task name",
	Long:  `Update the description of a specific # todo/task `,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		todoToRemove, err := strconv.ParseInt(args[0], 10, 64)

		if err != nil {
			fmt.Println("Please provide a valid todo number to update eg. 1")
			return
		}

		updatedTodoName := strings.ReplaceAll(
			strings.Join(args[1:], " "), string(TODO_FILE_DELIM), "")


		readTodosFromFile()
		refreshTodoDueDatesIfNeeded()

		var count int64 = 0
		found := false

		for i,todo := range todos {
			if !todo.completed {
				if count+1 == todoToRemove {
					todos[i].name = updatedTodoName
					found = true
					fmt.Printf("Todo #%d, has been udpated\n\n", todoToRemove)
					break
				}
				count += 1
			}
		}

		if !found {
			fmt.Printf("Could not find todo #%d to update.\n\n", todoToRemove)
			printTodos(false)
			return
		}

		printTodos(false)

		writeTodosToFile()
	},
}
