package main

import "fmt"

/*

chapter one is hello world chapter but there is some extra
if you use gofmt command it will format your file.
you can use:

gofmt -w filename.go
	to change file directly
gofmt -d filename.go
	to see what is not convenient to format
gofmt -l *
	to see which files is needed to format
gofmt filename.go
	to see formatted file content on stdout

note that gofmt does not format comment lines xd

*/

func main() {

	fmt.Println("chapter1")

}
