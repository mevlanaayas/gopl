package main

import (
	"flag"
	"fmt"
	"github.com/gopl/tempconv"
)

/*
we will convert command line arguments to the desired type (Celsius or Fahrenheit)
by using our package gopl/tempenv

we are using flags which we learnt in pointers.
user can call our function like:
go run import.go -v 200 -d fah
go run import.go -v -133 -d cel
*/


var value = flag.Float64("v", 100, "Convert given value to desired type")
var desired = flag.String("d", "cel", "cel stands for Celsius, fah stands for Fahrenheit")

func main()  {
	/*
	important statement stands here
	if you do not call the parse func flags will be set their default values
	when you call parse function.
	it parses the arguments and set flags to the given values
	*/
	flag.Parse()

	switch *desired {
	case "cel":
		fmt.Printf("In terms of %s value of %f equals to: %f\n", *desired, *value, tempconv.FtoC(tempconv.Fahrenheit(*value)))
	case "fah":
		fmt.Printf("In terms of %s value of %f equals to: %f\n", *desired, *value, tempconv.CtoF(tempconv.Celsius(*value)))

	}
}