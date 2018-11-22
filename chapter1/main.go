package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

const prefix string = "http://"

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

	/*
		about "nil"
		nil is a predeclared identifier representing the zero value for a
		pointer, channel, func, interface, map, or slice type.
		you can run: godoc builtin nil
	*/

	fmt.Println("chapter1")

	// runs a function according to the user input to handle un-compatible inputs for multiple functions
	switch os.Args[1] {
	case "echo":
		fmt.Println("*echo program*")
		echo(os.Args)
	case "unique":
		fmt.Println("*unique program with standard input*")
		result := make(map[string]int)
		unique(os.Stdin, result)
		fmt.Println("----- result -----", result)
	case "uniqueTwo":
		fmt.Println("*uniqueTwo program with standard/file input*")
		uniqueTwo()
	case "uniqueThree":
		fmt.Println("*uniqueThree program with standard/file (We use ReadFile) input*")
		uniqueThree()
	case "fetch":
		fmt.Println("*fetch program with standard input*")
		fetch()
	case "fetchTwo":
		fmt.Println("*fetchTwo program to print response to standard output*")
		fetchTwo()
	}

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

/*
we change unique func to handle both standard input and file input
note that standard input (os.Stdin) is also File pointer
also we will need to extend our results for multiple file case
so we pass results to function
*/
func unique(file *os.File, result map[string]int) {
	/*
		unique prints the text of each line that appears more than
		once in the text file (or standard input), preceded by its count
	*/
	/*
		create a map with key which is line
		lets loop over item and increase maps values by one
	*/
	args := bufio.NewScanner(file)

	for args.Scan() {
		if args.Text() == "exit" || args.Text() == string(0) || args.Text() == "" {
			break
		}
		result[args.Text()] += 1
	}
}

func uniqueTwo() {
	/*
		unique v2 looks for args from user. if user gives filename then
		runs uniques on file content
		runs on std in otherwise
		we must look os.Args to check user gives us a filename
	*/

	args := os.Args[1:]
	result := make(map[string]int)
	if len(args) == 0 {
		// there is no file name. time to take standard input
		fmt.Println("there is no filename found, calling unique()")
		unique(os.Stdin, result)
	}
	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			// we use print instead of panic to not break for loop.
			fmt.Println(err)
			continue
		}
		unique(file, result)
		file.Close()
	}
	fmt.Println("----- result -----", result)

}

func uniqueThree() {
	/*
		in this function we don't want to Open and read file manually (with Scanner)
		we want file's all contents at once.
		note that:
			we use ReadFile to do this job. And at last ReadFile function calls Open method ;)
			it is implemented to give choices to people
	*/
	args := os.Args
	result := make(map[string]int)
	if len(args) == 0 {
		fmt.Println("not collected any filename")
		unique(os.Stdin, result)
	}
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		/*
			we will have data which contains all data (included new lines)
			A big but here:
				return value of ReadFile is byte type. So type of data is byte.
			first we convert it to string then we split it from newlines
		*/
		for _, line := range strings.Split(string(data), "\n") {
			result[line] += 1
		}
	}
	fmt.Println("----- result -----", result)

}

func fetch() {
	/*
		take user input (from standard input)
		get urls (os.Args)
		and fetch all of them (http.Get)
	*/
	result := make(map[string]string)
	for _, url := range os.Args[2:] {
		if !strings.HasPrefix(url, prefix) {
			url = httpPatcher(url)
		}
		fmt.Println("url inside fetch function:", url)
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		response.Body.Close()
		result[url] = string(content)

	}
	fmt.Println("----- result -----", result)
}

func fetchTwo() {
	for _, url := range os.Args[2:] {
		if !strings.HasPrefix(url, prefix) {
			url = httpPatcher(url)
		}
		fmt.Println("url inside fetch function:", url)
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Response status:", response.Status)
		result, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("copy result:", result)
	}
}

func httpPatcher(url string) (fixedUrl string) {
	fixedUrl = prefix + url
	return
}
