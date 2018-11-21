package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

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

	/*
		about variable declaration
		we can use different forms
			exampleVariable := 10
			var exampleVariable int
			var exampleVariable int = 10
			var exampleVariable = 10
		so you can use any of them inside functions but
		for package-level you must use one starts with var
	*/

	fmt.Println("chapter1")

	for index, item := range os.Args {
		fmt.Println("at index:", index, ",arg:", item)
	}

	/*
		this is the simple echo example which prints arguments to the stdout
		two example about usage of args
		first garbage way
		second efficient way  (I think)

		I hope we can compare two methods by old way (by calculating execution times)

	*/
	// first method
	startTime := time.Now()
	var result, space string
	for _, item := range os.Args[1:] {
		result += space + item
		space = " "
	}
	fmt.Println("first method's result:", result)
	endTime := time.Now()
	firstDuration := endTime.Sub(startTime)
	fmt.Println("first duration:", firstDuration)

	//second method
	// we ignore first element ([1:]) because it is not an argument it is command itself
	startTime = time.Now()
	fmt.Println("second method's result:", strings.Join(os.Args[1:], " "))
	endTime = time.Now()
	secondDuration := endTime.Sub(startTime)
	fmt.Println("second duration:", secondDuration)

	// while loop
	condition := true
	count := 0
	for condition {
		// you can break the loop with "break" keyword for a certain condition
		if count == 5 {
			break
		}
		// you can break the loop by changing control condition
		if count == 10 {
			condition = false
		}
		// what happens if use return instead of break ???
		if count == 11 {
			return
		}
		count++
	}
	fmt.Println("count at the end of for loop:", count)

	arrayEx := [5]int{1, 2, 3, 4, 5}
	// we use reflect package to learn type of variable
	fmt.Println("what is the type of array:", reflect.TypeOf(arrayEx))
	for _, item := range arrayEx {
		fmt.Println(item)
	}

	/*
		note that we use underscore(_) to ignore assigning value.
		there is to option to ignore values.
		first is to create temp variable but golang does not allow unused vars.
		second one is to use _ underscore
		right side we have index and value of the element at that index
		but we only assign value to our item variable. and we ignore index (for this example).
	*/

	sliceEx := []int{1, 2, 3, 4, 5}
	fmt.Println("what is the type of slice:", reflect.TypeOf(sliceEx))
	for _, item := range sliceEx {
		fmt.Println(item)
	}

	/*
		we have string and every one character of is a byte.
		so when we iterate string we have a byte as an item.
	*/

	stringEx := "NewString"
	fmt.Println("what is the type of string:", reflect.TypeOf(stringEx))
	for _, item := range stringEx {
		fmt.Println("item as a byte:", item)
		fmt.Println("item as a string:", string(item))
	}

}
