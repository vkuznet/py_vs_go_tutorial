package main

import (
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", RequestHandler)
	http.ListenAndServe(":8888", nil)
}
