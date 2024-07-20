package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type ToDo struct {
	name      string
	completed bool
	created   time.Time
}

var todos = []ToDo{}
const TODO_FILE_DELIM = ','


func todoFilePath() string {
	dir, err := os.UserHomeDir()

	if err != nil {
		panic("User Home directory could not be found in $HOME")
	}

	var configDir string

	if dir[len(dir)-1] == '/' {
		configDir = dir + ".config"
	} else {
		configDir = dir + "/.config"
	}

	createDirectoryIfDoesNotExists(configDir)

	todoPath := configDir + "/gotodo"

	createDirectoryIfDoesNotExists(todoPath)

	return todoPath + "/todos.csv"
}

func createDirectoryIfDoesNotExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			panic("Could not create config dir " + path)
		}
   }
}

func writeTodosToFile() {
	var line string

	for _,todo := range todos {
		line += (todo.name + string(TODO_FILE_DELIM) +
			strconv.FormatBool(todo.completed) + string(TODO_FILE_DELIM) +
			todo.created.Format(time.RFC822Z) + "\n")
	}

	data := []byte(line)

	err := os.WriteFile(todoFilePath(), data, 0777)

	if err != nil {
		panic(err.Error())
	}
}

func readTodosFromFile() {
	file, err := os.ReadFile(todoFilePath())

	if err != nil {
		panic(err.Error())
	}

	l := 0
	// n := len(file)

	for i, b := range file {
		if b == '\n' {
			line := string(file[l:i])
			todoEntries := strings.Split(line, ",")

			if len(todoEntries) != 3 {
				panic("Invalid todo file format")
			}

			completed, err1 := strconv.ParseBool(todoEntries[1])
			dueDate, err2 := time.Parse(time.RFC822Z, todoEntries[2])

			if err1 != nil || err2 != nil {
				fmt.Println("Error parsing todo file, skipping")
			} else {
				todos = append(todos, ToDo{
					todoEntries[0],
					completed,
					dueDate,
				})
			}

			l = i + 1
		}
	}
}

func printTodos(completed bool) {
	fmt.Println("Todos")
	fmt.Println("-----")

	i := 1
	for _, td := range todos {
		if td.completed == completed {
			formattedTime := td.created.Format("Mon _2 3:04pm")
			fmt.Printf("%d. %s\n", i, td.name)
			fmt.Printf("  • %s\n", formattedTime)
			i++
		}
	}
}
