package main

import "fmt"

var global *int

func main() {
	/*
		--"new" function--
		we can also create a variable by using new builtin
		new(Type) create a new variable of type Type and returns its
		address.
	*/

	// following line is not true because new function returns an address so left side of assignment must be pointer
	//var nonPointerVar int = new(int)

	// note that *int is unnecessary here. because new(int) returns *int
	var pointerVar *int = new(int)
	var pointerVarTwo = new(int)
	fmt.Println(pointerVar)
	fmt.Println(pointerVarTwo)

	*pointerVar = 23
	fmt.Println(pointerVar)

	/*
		note that all print statements below prints addresses
		:) because they are pointer. to reach a value me must use "*".
		here we go!
	*/
	fmt.Println(*pointerVar)

	/*
		--"lifetime of variables"--
		package-level variables live through the program's runtime
		local variables live till they are unreachable

		garbage-collector : how it knows the variable life is ended?

		it is important to think about variables lifetime
		when you think about optimization.

	*/
	escapePossible()
	escapeNotPossible()
	fmt.Println(global)
	// again address :)
	fmt.Println(*global)

}

func escapePossible() {
	/*
		in this function local variable x lives even function ends.
		we call it x escapes from escapePossible
		it important when you think about performance optimizations
		because escapes requires extra memory allocation.
	*/
	var x int
	x = 1
	global = &x
}

func escapeNotPossible() {
	/*
		in this function when function called, y is created and no more lives when function finished.
	*/
	y := new(int)
	*y = 1
}
