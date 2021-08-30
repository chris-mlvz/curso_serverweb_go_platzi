package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start Server")
	server := NewServer(":3000")
	server.Handle("/", http.MethodGet, HandlerRoot)
	server.Handle("/create", http.MethodPost, PostRequest)
	server.Handle("/user", http.MethodPost, UserPostRequest)
	server.Handle("/api", http.MethodPost, server.AddMiddleWare(HandleHome, CheckAuth(), Loggin()))
	err := server.Listen()
	if err != nil {
		log.Fatal(err)
	}
}
