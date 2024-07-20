package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "a",
	Short: "Add a todo task",
	Long:  `Add a todo task to be reminded`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("No todo task specified")
			return
		}

		newTodo := strings.Join(args, " ")
		addTodoTask(strings.ReplaceAll(newTodo, string(TODO_FILE_DELIM), ""))
	},
}

func addTodoTask(name string) {
	readTodosFromFile()

	newTodo := ToDo{name, false, time.Now()}
	todos = append(todos, newTodo)

	writeTodosToFile()

	fmt.Printf("Added Todo: %s \n\n", name)
	printTodos(false)
}
