package main

// what we learn:
// bufio
import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}
	fmt.Println("Saving the note succeeded.")

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
