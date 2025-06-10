package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTest := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative", -10, false, "Negative is not prime, by definition"},
	}

	for _, e := range primeTest {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("\n%s: expect %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_promt(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	promt()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if string(out) != "-> " {
		t.Errorf("incorrect promt: expected -> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text not correct; got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", "Please enter a whole number!"},
		{"zero", "0", "0 is not prime, by definition!"},
		{"one", "1", "1 is not prime, by definition!"},
		{"two", "2", "2 is a prime number!"},
		{"three", "3", "3 is a prime number!"},
		{"negative", "-3", "Negative is not prime, by definition"},
		{"typed", "three", "Please enter a whole number!"},
		{"decimal", "30.2", "Please enter a whole number!"},
		{"quit", "q", ""},
		{"QUIT", "Q", ""},
		{"greek", "δεαθ", "Please enter a whole number!"},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
		}
	}
}

func Test_userInput(t *testing.T) {
	// To check this function we need a channel, and an instand of io.Reader
	doneChan := make(chan bool)

	// Create a reference to a bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)

}
