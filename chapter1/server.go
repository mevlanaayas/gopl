package main

import (
	"fmt"
	"net/http"
)

var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// we can also use Fprintf here instead of  w.Writer
	// w.Write([]byte(r.URL.Path))
	fmt.Println("--inside handler function--", r.URL.Path)
	/*
		when request comes to server it always call handler!!
		for /favicon.ico
		so we ignore other urls to guarantee exact count
	*/
	if r.URL.Path == "/favicon.ico" {
		count++
	}
	fmt.Fprintf(w, "Requested url is: %s", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--inside counter function--", r.URL.Path)
	fmt.Fprintf(w, "Current count is so far: %d", count)
}
