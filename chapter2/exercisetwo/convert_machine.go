package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
we will take any amount of numbers from command line.
if nothing comes from cli
we will look for standard input
	we will take till the user input "ok"

for each variable given by user. we will calculate:
FtoC CtoF // fahrenheit to celsius - celsius to fahrenheit
FtoM MtoF // feet to meter - meter to feet
PtoK KtoP // kilogram to pound- pound to kilogram
*/

/*
a good new: we will use interfaces (not the best case to use it but good to use it to practice)
instead of using six different types like Celsius, Fahrenheit, Meter...
we use a one CustomType
and interface to accept all the functions of custom type.
*/

type CustomType float64

func (f CustomType) CtoF() (c CustomType) {
	/*
		it is intentionally left default 5
	*/
	c = CustomType(5)
	return
}

func (f CustomType) FtoC() (c CustomType) {
	c = CustomType(5)
	return
}

func (f CustomType) FtoM() (c CustomType) {
	c = CustomType(5)
	return
}

func (f CustomType) MtoF() (c CustomType) {
	c = CustomType(5)
	return
}

func (f CustomType) KtoP() (c CustomType) {
	c = CustomType(5)
	return
}

func (f CustomType) PtoK() (c CustomType) {
	c = CustomType(5)
	return
}

// ------------------------
type Type interface {
	CtoF() CustomType
	FtoC() CustomType
	FtoM() CustomType
	MtoF() CustomType
	KtoP() CustomType
	PtoK() CustomType
}

func main() {
	/*
		for taken values we will create six different variables:
			Celsius(givenValue)
			Fahrenheit(givenValue)
			Feet(givenValue)
			Meter(givenValue)
			Kilogram(givenValue)
			Pound(givenValue)
		for new version we only need CustomType(givenValue)
	*/

	// TODO: we only support command line input. we need to implement standard input :)
	for _, arg := range os.Args[1:] {
		if givenValue, err := strconv.ParseFloat(arg, 64); err == nil {
			result(CustomType(givenValue))
		}
	}
}

func result(convertTypes ...Type) {
	for _, convertType := range convertTypes {
		fmt.Println("fahrenheit", convertType.CtoF())
		fmt.Println("celsius", convertType.FtoC())
		fmt.Println("meter", convertType.FtoM())
		fmt.Println("feet", convertType.MtoF())
		fmt.Println("pound", convertType.KtoP())
		fmt.Println("kilogram", convertType.PtoK())
		fmt.Println("-----------------------------")
	}
}
