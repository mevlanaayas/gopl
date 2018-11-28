package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() (str string) {
	/*
		string formatting
			check: "https://golang.org/pkg/fmt/"
	*/
	str = fmt.Sprintf("%g°C", c)
	return
}

func main() {
	/*
		you can use any of underlying type's properties with new type.
		we define new type Celsius over underlying type of float
		so we can use Celsius typed variable with float properties.
	*/
	fmt.Println("floats supports arithmetic so Celsius does:", BoilingC-FreezingC)
	BoilingF := CtoF(BoilingC)
	fmt.Println("same for Fahrenheit:", BoilingF)
	fmt.Println("same for Fahrenheit 2:", BoilingF-CtoF(FreezingC))

	/*
		but this will be compile error
			fmt.Println(BoilingC-BoilingF)
		BoilingC and BoilingF have different types (as expected)
	*/

	/*
		as being in arithmetic it is same for comparison
	*/
	var c Celsius
	var f Fahrenheit
	fmt.Println("celsius with 0:", c == 0)
	fmt.Println("fahrenheit with 0:", f >= 0)
	fmt.Println("fahrenheit with converted celsius:", f == Fahrenheit(c))

	/*
		note that we talked about "converted celsius"
		it means we converted type of celsius to type of fahrenheit.
		also note that type conversions change only only type of variables, not the values.
	*/

	/*
		lets try Celsius's String method
	*/
	newCels := FtoC(212)
	fmt.Println("expected °C:", newCels.String())
	fmt.Printf("when you format it calls String automatically according to the format type. %v\n", newCels)
	fmt.Printf("when you format it calls String automatically according to the format type. %v\n", newCels)
	fmt.Printf("expected float because format type is 'f': %f\n ", newCels)

}

/*
(c * 9/5 + 32)
*/
func CtoF(celsius Celsius) (fahrenheit Fahrenheit) {
	fahrenheit = Fahrenheit(celsius*9/5 + 32)
	return
}

/*
((f - 32) * 5 / 9)
*/
func FtoC(fahrenheit Fahrenheit) (celsius Celsius) {
	celsius = Celsius((fahrenheit - 32) * 5 / 9)
	return
}
