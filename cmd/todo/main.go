package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"todo"
)

const (
	todoFile = ".todos.json"
)

func main() {

	fmt.Print(`Commans for using: 
	-add : To add a new todo
	-complete number : To mark completed todo by number
	-delete number : Delete a todo by number
	-list : List all todos 
	-quit - To quit from the programm
	`)

	for {
		todos := &todo.Todos{}

		if err := todos.Load(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		var str string
		fmt.Scan(&str)
		switch {
		case str == "add":
			task, description, err := getInput(os.Stdin)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			todos.Add(task, description)
			err = todos.Store(todoFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		case str == "complete":
			fmt.Println("Enter the position of the mark: ")
			var position int
			fmt.Scan(&position)
			err := todos.Complete(position)

			if err != nil {
				fmt.Fprintln(os.Stdout, err)
				os.Exit(1)
			}

			todos.Store(todoFile)

		case str == "delete":
			fmt.Println("Enter the position of the mark: ")
			var position int
			fmt.Scan(&position)
			err := todos.Delete(position)

			if err != nil {
				fmt.Fprintln(os.Stdout, err)
				os.Exit(1)
			}

			todos.Store(todoFile)
		case str == "list":
			todos.Print()
		case str == "quit":
			fmt.Println("Quiting from the programm...")
			os.Exit(1)
		default:
			fmt.Fprintln(os.Stdout, "invalid command")
		}
	}
}

func getInput(r io.Reader) (string, string, error) {
	reader := bufio.NewReader(r)

	fmt.Print("Enter task: ")
	task, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter description: ")
	description, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	// Удаление символов новой строки из task и description
	task = strings.TrimSuffix(task, "\n")
	description = strings.TrimSuffix(description, "\n")

	return task, description, nil
}
