package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Task        string
	Description string
	Time        time.Time
	Done        bool
}

type Todos []item

func (t *Todos) Add(task string, description string) {
	item := item{
		Task:        task,
		Description: description,
		Time:        time.Now(),
		Done:        false,
	}

	*t = append(*t, item)
}

func (t *Todos) Complete(index int) error {

	if index < 0 || index > len(*t) {
		return errors.New("invalid index")
	}

	(*t)[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	list := *t

	if index < 0 || index > len(*t) {
		return errors.New("invalid index")
	}

	*t = append(list[:index-1], list[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)

	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)

	if err != nil {
		return nil
	}

	var f *os.File
	_, statErr := os.Stat(filename)
	if os.IsNotExist(statErr) {
		f, err = os.Create(filename)

		if err != nil {
			return err
		}
	} else {
		f, err = os.Open(filename)

		if err != nil {
			return err
		}
	}

	defer f.Close()

	os.WriteFile(filename, data, 0644)

	return nil
}

func (t *Todos) Print() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Description"},
			{Align: simpletable.AlignRight, Text: "Created at"},
			{Align: simpletable.AlignRight, Text: "Done"},
		},
	}

	var cells [][]*simpletable.Cell

	for i, item := range *t {
		i++
		task := blue(item.Task)
		done := blue("no")
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("yes")

		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", i)},
			{Text: task},
			{Text: item.Description},
			{Text: item.Time.Format(time.RFC1123)},
			{Text: done},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d uncompleted tasks", t.CountTasks()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)
	table.Print()
}

func (t *Todos) CountTasks() int {
	result := 0

	for _, item := range *t {
		if !item.Done {
			result++
		}
	}

	return result
}
