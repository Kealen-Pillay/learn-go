package main

import "fmt"

func main() {
	foo := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range foo {
		if foo[i] % 2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}