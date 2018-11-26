package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	/*

	 */
	var variable = 10
	var addressOfVariable = &variable

	var pointerToVariable *int
	pointerToVariable = &variable
	fmt.Println(reflect.TypeOf(addressOfVariable), reflect.TypeOf(pointerToVariable))
	fmt.Println(*addressOfVariable)

	// two variables with zero values
	zeroVariable := 0
	var uninitializedVariable int

	// if pointer points to a variable, pointer != nil, will be true
	if &zeroVariable != nil {
		fmt.Println("Y pointer is not nil")
	}
	if &uninitializedVariable != nil {
		fmt.Println("T pointer is not nil")
	}

	/*
		basic property of pointers.
		we can send them into function or indirectly reach their values
	*/
	var returnValue = f()

	fmt.Println("value 1, expected address", returnValue)
	fmt.Println("value 2, expected value", *returnValue)

	/*
		every function call create new variable so the address of these variables
		will be different, so the return value will be different.
	*/
	fmt.Println("expected false", f() == f())

	/*

	 */
	variableV := 1
	incr(&variableV) // this will change the value even we did not assign result
	fmt.Println("two times incremented variable", incr(&variableV))

	/*

	 */
	variableVTwo := incr(&variableV)
	fmt.Println("final Variable 2", variableVTwo)
	fmt.Println("final Variable", variableV)

	/*
		we use the flag package to use flags from cli
		n to insert new line at the end
		sep to separate line with different separator instead of " "
	*/
	echoFour()
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func echoFour() {
	/*
		note that
			n and sep are pointers. flag.Type returns pointer of type Type
	*/
	// expecting address
	fmt.Println("address of n", n)
	fmt.Println("address of sep", sep)
	// expecting values
	fmt.Println("value of n", *n)
	fmt.Println("value of sep", *sep)

	fmt.Println("os.args before parse", os.Args)
	fmt.Println("args before parse", flag.Args())

	/*
		flag.Parse()
		simply loop over all os.Args, dispatch the values with "-" in front of them which is flag,
		returns the cleaned arguments
	*/
	flag.Parse()
	fmt.Println("os.args after parse", os.Args)
	fmt.Println("args after parse", flag.Args())
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

}
