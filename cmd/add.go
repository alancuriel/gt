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

	newTodo := ToDo{name, false, nextEndOfWorkDay()}
	todos = append(todos, newTodo)

	writeTodosToFile()

	fmt.Printf("Added Todo: %s \n\n", name)
	printTodos(false)
}

func nextEndOfWorkDay() time.Time {
	now := time.Now()
	currentWeekday := now.Weekday()

	if currentWeekday == time.Saturday {
		now = now.AddDate(0, 0, 2)
	} else if currentWeekday == time.Sunday {
		now = now.AddDate(0, 0, 1)
	}

	eod := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location())

	now2 := time.Now()

	if eod.After(now2) {
		return eod
	}


	if eod.Weekday() == time.Friday {
		eod = eod.AddDate(0, 0, 3)
	} else {
		eod = eod.AddDate(0, 0, 1)
	}

	return eod
}
