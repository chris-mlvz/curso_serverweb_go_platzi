package main

import (
	"fmt"
	"net/http"
)

func main() {
	// start := time.Now()
	ch := make(chan string, 4)

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for {
		for _, server := range servers {
			go checkServer(server, ch)
		}
		select {
		case v := <-ch:
			fmt.Println(v)
		}
	}

	// elapsed := time.Since(start)
	// fmt.Println("Time elapsed:", elapsed)
}

func checkServer(server string, ch chan<- string) {
	_, err := http.Get(server)
	if err != nil {
		ch <- server + " is not available"
		return
	}
	ch <- server + " is working normally"
}
