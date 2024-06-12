package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnglishBot struct {

}

type SpanishBot struct {

}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
 func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
 }

// Can omit receiver type name is we are not using the struct
func (EnglishBot) getGreeting() string {
	// custom logic for eb
	return "Wagwan"
}

func (SpanishBot) getGreeting() string {
	// custom logic for sb
	return "Te Amo"
}