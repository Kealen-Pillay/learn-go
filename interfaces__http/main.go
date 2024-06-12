package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type LogWriter struct {}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	lw := LogWriter{}
	io.Copy(lw, resp.Body)
}

func (LogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote out %v bytes", len(bs))
	return len(bs), nil
}