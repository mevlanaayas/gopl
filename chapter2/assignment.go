package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	/*
		= sign
		it assign right of sign to left of sign
	*/

	// declarations to use in example
	assignExampleVar := 10
	var pointer = new(int)
	var mevlana = Person{"mevlana", 25}
	var array = [5]int{5, 6, 7, 8, 9}

	// example assignments
	assignExampleVar = 20
	*pointer = 30
	mevlana.name = "mevlana ayas"
	array[3] = array[2] + 6

	// to prevent unused variable error
	fmt.Println(assignExampleVar)

	// arithmetic and bitwise operators have a short assignment
	/*
	bitwise operators
	| or
	& and
	^ xor
	<< left shift
	>> right shift
	*/
	assignExampleVar += 5 // which means assignExampleVar = assignExampleVar + 5

	assignExampleVar |= 3

	// also increment and decrement can be done by
	assignExampleVar++
	assignExampleVar--

}
