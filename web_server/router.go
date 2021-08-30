package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path, meth string) (http.HandlerFunc, bool, bool) {
	_, pathExist := r.rules[path]
	handler, methExist := r.rules[path][meth]
	return handler, pathExist, methExist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, pathExist, methExist := r.FindHandler(request.URL.Path, request.Method)
	if !pathExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !methExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}
