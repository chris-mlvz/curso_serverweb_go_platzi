package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(nextHandler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, request *http.Request) {
			flag := true
			fmt.Println("Checking Authentication")
			if flag {
				nextHandler(w, request)
			} else {
				return
			}
		}
	}
}

func Loggin() Middleware {
	return func(nextHandler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			nextHandler(w, r)
		}
	}
}
