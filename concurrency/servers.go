package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}
	for _, server := range servers {
		wg.Add(1)
		go checkServer(server)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Time elapsed:", elapsed)
}

func checkServer(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "is not available")
		wg.Done()
	} else {
		fmt.Println(server, "is working normally")
		wg.Done()
	}
}
