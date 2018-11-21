package main

import (
	"bufio"
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

	fmt.Println("*echo program*")
	echo(os.Args)

	fmt.Println("*unique program*")
	unique()

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

	/*
		formatted print function
		* It does not create new line at the end use \n
		for tab use \t
		* You need to specify type of variable
		IMO It is more comfortable to use variables inside printed string
		for example:
	*/

	numberOne := 0
	numberTwo := 1
	infoText := "We use %s"
	fmt.Printf("\t number one: %d, and number two %d \n", numberOne, numberTwo)
	fmt.Printf("For string %s \n", infoText)
	fmt.Printf(
		"sum of %d and %d is equals to: %f (converted to float)\n",
		numberOne,
		numberTwo,
		float64(numberOne+numberTwo),
	)

}

func echo(args []string) {
	for index, item := range args {
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
	for _, item := range args[1:] {
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
	fmt.Println("second method's result:", strings.Join(args[1:], " "))
	endTime = time.Now()
	secondDuration := endTime.Sub(startTime)
	fmt.Println("second duration:", secondDuration)

}

func unique() {
	/*
		unique prints the text of each line that appears more than
		once in the text file (or standard input), preceded by its count
	*/
	/*
		create a map with key which is line
		lets loop over item and increase maps values by one
	*/
	args := bufio.NewScanner(os.Stdin)
	result := make(map[string]int)

	for args.Scan() {
		if args.Text() == "exit" || args.Text() == string(0) || args.Text() == "" {
			break
		}
		result[args.Text()] += 1
	}
	fmt.Println("----- result -----", result)
}
