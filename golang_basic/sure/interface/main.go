package main

// what we learn:
// bufio, any/interface{} type, generic type
import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputable interface {
	saver
	displayer
}

func main() {
	title, content, err := getNoteData()

	todoText, _ := getUserInput("Input todo data: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	output(userNote)
	output(todo)
}

func output(data outputable) {
	data.Display()
	data.Save()
}
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving succeeded.")
	return nil
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note Title:")
	if err != nil {
		return "", "", err
	}
	content, err := getUserInput("Note Content:")
	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Printf("%v ", prompt)
	// var value string
	// fmt.Scan(&value)
	// if value == "" {
	// 	return "", errors.New("Invalid input.")
	// }
	// return value, nil
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	text = strings.TrimSuffix(text, "\n")

	return text, nil
}
