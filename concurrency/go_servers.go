package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string, 4)

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}
	for _, server := range servers {
		go checkServer(server, ch)
	}

	for range servers {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Println("Time elapsed:", elapsed)
}

func checkServer(server string, ch chan<- string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "is not available")
		ch <- "error"
		return
	}
	fmt.Println(server, "is working normally")
	ch <- "ok"
}
