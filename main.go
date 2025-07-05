package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"test/note"
	"test/todo"
)

func main() {
	title, content := getNoteData()
	newNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	newNote.Display()

	if err = newNote.Save(); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Saving the note succeeded")
	newTodo, err := todo.New(getUserInput("Enter todo text: "))

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	newTodo.Display()

	if err = newTodo.Save(); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Saving the todo succeeded")
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
