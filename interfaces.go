package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
	// getGreeting(string, int) (string, error)
	//  function    arguments     return type
	//
	// getBotVersion() float64
	// respondToUser(user) string
}

// Concrete types:
// - englishBot
// - map
// - struct
// - int
// - string

// Interface type
// - bot

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	// custom logic to generate english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// custom logic to generate spanish greeting
	return "Hoal!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
