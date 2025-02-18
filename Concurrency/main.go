package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {
	// go greet("nice to meet you!")
	// go greet("nice to meet you!")
	done := make(chan bool)
	go slowGreet("How...Are...You", done)
	go greet("nice to meet you!")
	<-done
}
