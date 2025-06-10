package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()

	doneChan := make(chan bool)

	go readUserInput(os.Stdin, doneChan)

	<-doneChan

	close(doneChan)

	fmt.Println("Goodbye.")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)
	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		promt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

func intro() {
	fmt.Println("Is it prime?")
	fmt.Println("____________")
	fmt.Println("Enter a whole number, and we'll tell you if it is prime or not. Enter q to quit.")
	promt()
}

func promt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	// 0 and 1 is not prime by definition
	if n == 1 || n == 0 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative is not prime by definition
	if n < 0 {
		return false, "Negative is not prime, by definition"
	}

	// use the modules operator repeatedly to see if we are have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d!", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
