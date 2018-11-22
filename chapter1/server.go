package main

import "net/http"

func main()  {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", hello)
	go func() {
		http.ListenAndServe("localhost:8080", nil)
	}()
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello There"))
}

func hello(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello World"))
}