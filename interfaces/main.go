package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	// eb := englishBot{}
	// sb := spanishBot{}

	// printGreeting(eb)
	// printGreeting(sb)

	resp := makeRequest("https://google.com")

	if resp != "" {
		fmt.Println("Request was successful", resp)
	} else {
		println("Request failed")
	}

}

func printGreeting(b bot) {
	println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
