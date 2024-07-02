package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"todo"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark completed todo")
	delete := flag.Int("delete", 0, "delete a todo")
	list := flag.Bool("list", false, "List all todos")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task, description, err := getInput(os.Stdin, flag.Args()...)
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
	case *complete > 0:
		err := todos.Complete(*complete)

		if err != nil {
			fmt.Fprintln(os.Stdout, err)
			os.Exit(1)
		}

		todos.Store(todoFile)

	case *delete > 0:
		err := todos.Delete(*delete)

		if err != nil {
			fmt.Fprintln(os.Stdout, err)
			os.Exit(1)
		}

		todos.Store(todoFile)
	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stdout, "invalid comand")
		os.Exit(1)
	}
}

func getInput(r io.Reader, args ...string) (string, string, error) {

	var task, description string

	fmt.Print("Enter task: ")
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", "", nil
	}

	task = scanner.Text()

	fmt.Print("Enter description: ")
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", "", nil
	}
	description = scanner.Text()

	if len(task) == 0 {
		return "", "", errors.New("empty todo")
	}

	return task, description, nil
}
