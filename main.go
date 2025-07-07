package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"test/note"
	"test/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	newNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	if err = outputData(newNote); err != nil {
		return
	}
	newTodo, err := todo.New(getUserInput("Enter todo text: "))

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_ = outputData(newTodo)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	if err := data.Save(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("Saving the \"%s\" succeeded\n", reflect.TypeOf(data).String())

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	value, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return ""
	}
	value = strings.TrimSuffix(strings.TrimSuffix(value, "\n"), "\r")

	return value
}
