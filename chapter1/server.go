package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// we can also use Fprintf here instead of  w.Writer
	// w.Write([]byte(r.URL.Path))
	fmt.Fprintf(w, "Requested url is: %s", r.URL.Path)
}
