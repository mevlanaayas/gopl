package main

import (
	"fmt"
	"os"
)

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

	/*
		tuple assignments
		for assigning multiple variables at once.
		right-hand side of assign operator evaluated before any left-hand side variables updated.
	*/

	// variable declarations to use them for assignment
	// note that we use here tuple assignment also :)
	channelVar := make(chan int)
	go func() {
		channelVar <- 2
		channelVar <- 3
	}()
	var typeAssertion interface{} = "example type assertion"
	mapVar := make(map[int]int)
	first, second, third := 1, 3, 10
	var swappableArray = [4]int{1, 2, 3, 4}
	key := 10

	// e.g swapping
	first, second = second, first
	// to prevent unused variable error
	fmt.Println(first, second, third)
	swappableArray[1], swappableArray[2] = swappableArray[0], swappableArray[3]

	fmt.Println(gcd(6, 24))
	fmt.Println(fib(7))

	/*
		tuple assignment allow trivial assignments
		but make sure your statements not complex.
		separate statements is easier to read
	*/
	first, second, third = gcd(first, second), third, fib(third) // instead of this use separate state

	/*
		functions can return multiple result,  produce more than one variable.
		we must sure left hand side always fit with function's return value.
		examples:
		file opening
		map lookup
		type assertion
		channel

	*/
	f, err := os.Open("input.txt")
	fmt.Println("file:", f, err)
	v, ok := mapVar[key]
	fmt.Println("map:", v, ok)
	str, ok := typeAssertion.(string)
	fmt.Println("assert str:", str, ok)
	integer, ok := typeAssertion.(int)
	fmt.Println("assert int:", integer, ok)
	flt, ok := typeAssertion.(float64)
	fmt.Println("assert float:", flt, ok)
	com, ok := <-channelVar
	fmt.Println("channel:", com, ok)

	/*
		note that we do not have to assign all returning values to the variables
		we can ignore some of returning values. (if we ignore all statement will be unnecessary xd)
		just use "_" underscore to ignore
	*/
	result, _ := <-channelVar // we ignore the variable which holds error state
	fmt.Println("result:", result)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(x int) int {
	/*
		iterative version of fibonacci sequence
	*/
	m, n := 0, 1
	for i := 0; i < x; i++ {
		m, n = n, m+n
	}
	return m
}
