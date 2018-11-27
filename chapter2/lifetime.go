package main

import "fmt"

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

}
