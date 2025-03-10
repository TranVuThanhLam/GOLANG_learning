package cmdmanager

import "fmt"

type CmdManager struct{}

func (cmd CmdManager) ReadLines() ([]string, error) {
	fmt.Println("Readlines result: ")
	return nil, nil
}

func (cmd CmdManager) WriteResult(data any) error {
	fmt.Println("Write result:")
	return nil
}
