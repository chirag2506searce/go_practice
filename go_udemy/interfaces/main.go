package main

import (
	"fmt"
)

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string // every 'type' that has a function getGreeting() string, becomes a bot
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)
	printGreeting(sb)
}

// func printGreeting(eb englishBot) {
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

func (englishBot) getGreeting() string {
	// custom logic for english greeting (just hello for demonstration)
	return "Hello!" /*this could be kept as field while defining struct, but the motive is to show
	that both structs have similar function*/
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
