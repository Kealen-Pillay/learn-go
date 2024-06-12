package main

import (
	"fmt"
	"net/http"
	"time"
)

// Concurrency
// ===========
// We can have multiple threads executing code. If one thread blocks another one is picked upa nd worked on

// Parallelism
// ===========
// Multiple threads executed at the exact same time. Requires multiple CPU's

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, url := range urls {
		go checkLink(url, c)	//go routine
	}

	for l := range c {
		time.Sleep(5 * time.Second)
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}
	fmt.Println(link, " is up")
	c <- link
}