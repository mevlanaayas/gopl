package main

import "fmt"

func main() {
	/*
		var name type = expression
		either type or "= expression" may be missed // type is actually unnecessary there
		such as:
			var name = {value}
			var name int
		in the second way variable has a default(zero) value according to the type.
			0 for numbers,
			false for booleans,
			"" for strings,
			nil for interfaces and reference types (slice, pointer, map, channel, function).
			The zero value of array or a struct has the zero value of all of its elements or fields.
	*/
	var variableName int = 5
	var variableWithoutExpression string
	fmt.Println(variableName, variableWithoutExpression)

	var variableWithoutType = "asd"
	fmt.Println(variableWithoutType)

	/*
		it possible multiple variable declaration
	*/
	var i, j int
	var b, f = 2.3, "four"
	fmt.Println(i, j, b, f)

	/*
		short variable declaration
	*/
	zeroPoint := 0
	count, item := 10, "item"
	fmt.Println(zeroPoint, count, item)

	/*
		swap values in single line
	*/
	i, j = count, variableName

	/*
		short variable declaration behave as assignment
		when you use variable which is already defined, with multiple short variable declaration statement
		it behaves as assignment operator!
		A big but here
			for this rule at least one variable must be undeclared before.
	*/
	existVariable, secondExistVariable := 10, 10
	newVariable, existVariable := 20, 12
	fmt.Println(existVariable, newVariable, secondExistVariable)

	// this will throw an error.
	// existVariable, secondExistVariable := 1, 1

}
