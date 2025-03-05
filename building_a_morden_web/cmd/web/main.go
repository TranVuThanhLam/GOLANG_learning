package main

import (
	"fmt"
	"net/http"

	"modern_web.com/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	// http.HandleFunc("/about", handlers.About)

	fmt.Printf("server is running on port%s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
