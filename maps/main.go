package main

import "fmt"

func main() {

	colors := map[string]string{
		"red": "#wewe23",
		"green": "#asdasd3",
	}
	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color, hex)
	}
}